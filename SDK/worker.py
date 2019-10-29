import os, time, json
from .const.const import wt, dt, ds
from .utilslib.logger import logger
from .utilslib.api import api
from .utilslib.fileops import load_json_file

def get_environment():
    debug = os.environ.get('DEBUG', False) #是否打印调试信息
    rootdir = os.environ.get('ROOTDIR', 'scratch/') #数据的根目录
    apihost = os.environ.get('WEBURL', 'http://192.168.1.100:9000')
    return debug, rootdir, apihost

class worker_fs():
    def __init__(self):
        #一级目录
        self.csv_dir = os.path.join(self.rootdir, 'csv')
        self.img_dir = os.path.join(self.rootdir, 'img')
        self.datasets_dir = os.path.join(self.rootdir, 'datasets')
        self.projects_dir = os.path.join(self.rootdir, 'projects')
        self.scratch_dir  = os.path.join(self.rootdir, 'scratch')
        self.modules_dir  = os.path.join(self.rootdir, 'modules')

        if self.wtype == wt.DATA.value:
            self.dataset_dir = os.path.join(self.datasets_dir, self.wdir)
            self.logfile = os.path.join(self.dataset_dir, 'log.txt')
            self.info_json = os.path.join(self.dataset_dir, 'info.json')
            self.info2_json = os.path.join(self.dataset_dir, 'info2.json')
            self.dataset_lists = os.path.join(self.dataset_dir, 'filelist.csv')
            self.dataset_cellslists = os.path.join(self.dataset_dir, 'cellslist.csv')
            self.cells_dir = os.path.join(self.dataset_dir, 'cells')
            self.cells_crop_dir = os.path.join(self.cells_dir, 'crop')
            self.cells_crop_masked_dir = os.path.join(self.cells_dir, 'crop_masked')
            self.cells_statistics_dir = os.path.join(self.cells_dir, 'statistics')
        elif self.wtype == wt.TRAIN.value:
            self.project_dir = os.path.join(self.projects_dir, self.wdir)
            self.info_json = os.path.join(self.project_dir, 'info.json')
            self.project_train_dir = os.path.join(self.project_dir, 'train')
            self.project_resize_train_dir = os.path.join(self.project_dir, 'resize_train')
            self.project_test_dir = os.path.join(self.project_dir, 'test')
            self.project_resize_test_dir = os.path.join(self.project_dir, 'resize_test')
            self.project_train_labels_csv = os.path.join(self.project_dir, 'train_labels.csv')
            self.project_test_labels_csv = os.path.join(self.project_dir, 'test_labels.csv')
            self.project_mod_path = os.path.join(self.project_dir, 'Modelak.h5')
            self.project_tmp_dir = os.path.join(self.project_dir, 'autokeras')
        elif self.wtype == wt.PREDICT.value:
            self.project_dir = os.path.join(self.projects_dir, self.wdir)
            self.info_json = os.path.join(self.project_dir, 'info.json')
            self.project_predict_dir = os.path.join(self.project_dir, 'predict')
            self.project_resize_predict_dir = os.path.join(self.project_dir, 'resize_predict')
            self.project_predict_error_data_dir = os.path.join(self.project_dir, 'predict_error_data')
            self.project_predict_labels_csv = os.path.join(self.project_dir, 'predict_labels.csv')
        else:
            raise RuntimeError("unknown worker type %d" % self.wtype)

    def checkdir(self):
        dirs = [self.csv_dir, self.img_dir]
        for folder in dirs:
            if not os.path.exists(folder):
                raise RuntimeError("not found %s of %d" % (folder, self.wid))
        #创建处理数据集的目录
        cells_dirs = []
        worker_dir = None
        if self.wtype == wt.DATA.value:
            worker_dir = self.dataset_dir
            cells_dirs = [self.cells_dir, self.cells_crop_dir, self.cells_crop_masked_dir, self.cells_statistics_dir]
        elif self.wtype == wt.TRAIN.value:
            worker_dir = self.project_dir
            cells_dirs = [self.project_train_dir]
        elif self.wtype == wt.PREDICT.value:
            worker_dir = self.project_dir
            cells_dirs = [self.project_predict_dir]

        if not os.path.exists(self.project_dir):
            raise RuntimeError("not found %s of %d" % (self.dataset_dir, self.wid))
        for folder in cells_dirs:
            if not os.path.exists(folder):
                os.makedirs(folder)

    def load_info_json(self):
        info = load_json_file(self.info_json)

        #训练或者预测时候找到被训练或者被预测的数据集，细胞信息全在cellslist.csv文件
        if self.wtype == wt.TRAIN.value or self.wtype == wt.PREDICT.value:
            self.cellslist_csv =  self.cellslist_path(info['ddir'])
        return info

    def save_info_json(self, info, savefile):
        with open(savefile, 'w', encoding='utf-8') as file:
            file.write(json.dumps(info, indent=2, ensure_ascii=False))
        return

    def cells_cache_path(self, batchid, medicalid, cellname):
        return os.path.join(self.scratch_dir, batchid, medicalid, cellname)

    #传入任意数据集的dirname，返回这个目录下的cellslist.csv路径
    def cellslist_path(self, wdir):
        return os.path.join(self.datasets_dir, wdir, 'cellslist.csv')

class worker_api(api):
    def __init__(self, apihost, wtype):
        self.apihost = apihost
        self.wtype = wtype
        api.__init__(self, self.apihost, self.debug)

    def get_job(self):
        wid, dirname = 0, None
        _status = ds.READY4PROCESS.value

        wid, status, dirname = self.get_one_job(_status, self.wtype)
        #检查得到的任务是不是想要的
        if status != _status or dirname is None:
            wid, dirname = 0, None
        return wid, dirname
    def woker_percent(self, percent, ETA):
        self.percent = percent
        self.post_job_status(self.wid, self.status, self.wtype, self.percent, ETA)

class worker(worker_api, worker_fs):
    def __init__(self, wtype):
        self.wtype = wtype    #任务的类型
        self.debug, self.rootdir, self.apihost = get_environment()
        self.percent, self.status = 0, 0
        self.ETA = 0 #预估还要多长时间结束

        #API初始化
        worker_api.__init__(self, self.apihost, self.wtype)

        #日志
        self.logfile = None
        self.log = logger(self.logfile)

        #初始化完成
        self.log.info("worker initialized %d" % time.time())

    def prepare(self, workerid, workerdir, wtype):
        self.wid = workerid   #任务id，处理数据表示数据集id，训练/预测表示项目id
        self.wdir = workerdir #任务执行的目录，处理数据表示数据集目录，训练/预测表示项目目录

        #目录和文件初始化
        worker_fs.__init__(self)
        self.checkdir()

        #初始化日志到工作目录
        self.log = logger(self.logfile)

        self.log.info("worker run wid=%d wdir=%s wtype=%d" % (workerid, workerdir, wtype))
        self.percent = 0
        self.status = ds.PROCESSING.value

    def done(self):
        self.status = ds.PROCESSING_DONE.value
        self.woker_percent(100, 0)

    def error(self):
        self.status = ds.PROCESSING_ERR.value
        self.woker_percent(self.percent, self.ETA)
