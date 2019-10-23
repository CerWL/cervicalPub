import os, time
from const.const import wt, dt, ds
from utilslib.logger import logger
from utilslib.api import api

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

        if self.wtype == wt.DATA.value:
            self.dataset_dir = os.path.join(self.datasets_dir, self.wdir)
            self.dataset_json = os.path.join(self.dataset_dir, 'dataset.json')
            self.cells_dir = os.path.join(self.dataset_dir, 'cells')
            self.cells_crop_dir = os.path.join(self.cells_dir, 'crop')
            self.cells_crop_masked_dir = os.path.join(self.cells_dir, 'crop_masked')
            self.cells_statistics_dir = os.path.join(self.cells_dir, 'statistics')
        elif self.wtype == wt.TRAIN.value:
            self.project_dir = os.path.join(self.projects_dir, self.wdir)
        elif self.wtype == wt.PREDICT.value:
            self.project_dir = os.path.join(self.projects_dir, self.wdir)
        else:
            raise RuntimeError("unknown worker type %d" % self.wtype)
    def checkdir(self):
        dirs = [self.csv_dir, self.img_dir]
        for folder in dirs:
            if not os.path.exists(folder):
                raise RuntimeError("not found %s of %d" % (folder, self.wid))
        #创建处理数据集的目录
        if self.wtype == wt.DATA.value:
            if not os.path.exists(self.dataset_dir):
                raise RuntimeError("not found %s of %d" % (self.dataset_dir, self.wid))
            cells_dirs = [self.cells_dir, self.cells_crop_dir, self.cells_crop_masked_dir, self.cells_statistics_dir]
            for folder in cells_dirs:
                if not os.path.exists(folder):
                    os.makedirs(folder)

class worker_api(api):
    def __init__(self, apihost, wtype):
        self.apihost = apihost
        self.wtype = wtype
        api.__init__(self, self.apihost, self.debug)

    def get_job(self):
        wid, dirname = 0, None
        _status = ds.READY4PROCESS.value
        if self.wtype == wt.TRAIN.value:
            _status = ds.READY4TRAIN.value
        elif self.wtype == wt.PREDICT.value:
            _status = ds.READY4PREDICT.value

        wid, status, dirname, jobtype = self.get_one_job(_status, self.wtype)
        #检查得到的任务是不是想要的
        if status != _status or dirname is None:
            wid, dirname = 0, None
        return wid, dirname

class worker(worker_fs, worker_api):
    def __init__(self, workerid, workerdir, wtype):
        self.wid = workerid   #任务id，处理数据表示数据集id，训练/预测表示项目id
        self.wdir = workerdir #任务执行的目录，处理数据表示数据集目录，训练/预测表示项目目录
        self.wtype = wtype    #任务的类型
        self.debug, self.rootdir, self.apihost = get_environment()

        #目录和文件初始化
        worker_fs.__init__(self)
        self.checkdir()
        #API初始化
        worker_api.__init__(self, self.apihost, self.wtype)

        #日志
        self.logfile = os.path.join(self.rootdir, 'log.txt')
        self.log = logger(self.logfile, stdout=self.debug)

        #初始化完成
        self.log.info(self.projects_dir)
        self.log.info("worker initialized %d" % time.time())

w = worker(1, "11", 1)
wid, wdir = w.get_job()
w.log.warning("%s %s" % (wid, wdir))
