# -*- coding: utf-8 -*-
import os, csv, cv2, time, argparse, gc
from enum import Enum
from autokeras.utils import pickle_from_file
from autokeras.image.image_supervised import load_image_dataset, ImageClassifier
from keras.preprocessing.image import load_img, img_to_array
import numpy as np
from shutil import copyfile, rmtree
import pandas as pd
from utilslib.jsonfile import load_json_file, update_model_info_json
from utilslib.webserverapi import get_one_job, post_job_status

localdebug = os.environ.get('DEBUG', 'False')

# datasets status
class ds(Enum):
    INIT            = 0 #0初始化
    READY4PROCESS   = 1 #1用户要求开始处理
    PROCESSING      = 2 #2开始处理
    PROCESSING_ERR  = 3 #3处理出错
    PROCESSING_DONE = 4 #4处理完成
    PATH_ERR        = 5 #5目录不存在
    READY4TRAIN     = 6 #6用户要求开始训练
    TRAINNING       = 7 #7开始训练
    TRAINNING_ERR   = 8 #8训练出错
    TRAINNING_DONE  = 9 #9训练完成

# datasets type
class dt(Enum):
    UNKNOWN  = 0 #0未知
    TRAIN    = 1 #1训练
    PREDICT  = 2 #2预测

def timestamp():
    return time.strftime("%Y%m%d_%H%M%S", time.localtime())

#传入文件的路径，返回路径，文件名字，文件后缀
def get_filePath_fileName_fileExt(filename):
    (filepath, tempfilename) = os.path.split(filename)
    (shotname, extension) = os.path.splitext(tempfilename)
    return filepath, shotname, extension

#write csv
def write_csv(img_dir, csv_dir):
    list = []
    list.append(['File Name','Label'])
    for file_name in os.listdir(img_dir):
    	for img in os.listdir("%s/%s"%(img_dir,file_name)):
            #print (img)
            item = [file_name+"/"+img, file_name]
            list.append(item)
    f = open(csv_dir, 'w')
    writer = csv.writer(f)
    writer.writerows(list)

#resize images
def resize_img(input_dir, output_dir, RESIZE):
    cls_file = os.listdir(input_dir)
    for cls_name in cls_file:
        img_file = os.listdir("%s/%s"%(input_dir,cls_name))
        for img_name in img_file:
            #print (img_name)
            img = cv2.imread("%s/%s/%s"%(input_dir,cls_name,img_name))
            if img.shape[0] != img.shape[1]:
                print("skip this image w != h: %s" % img_name)
                continue
            img = cv2.resize(img,(RESIZE,RESIZE),interpolation=cv2.INTER_LINEAR)
            if os.path.exists("%s/%s"%(output_dir,cls_name)):
                cv2.imwrite("%s/%s/%s"%(output_dir,cls_name,img_name),img)
            else:
                os.makedirs("%s/%s"%(output_dir,cls_name))
                cv2.imwrite("%s/%s/%s"%(output_dir,cls_name,img_name),img)
#resize images
def resize_img2(typetree, output_dir, RESIZE):
    for i in typetree.keys():
        for j in typetree[i]:
            filepath, shotname, extension = get_filePath_fileName_fileExt(j)
            typedir = os.path.join(output_dir, i)
            if not os.path.exists(typedir):
                os.makedirs(typedir)
            img = cv2.imread(j)
            img_name = shotname + extension
            if img.shape[0] != img.shape[1]:
                print("skip this image w != h: %s" % img_name)
                continue
            img = cv2.resize(img, (RESIZE, RESIZE), interpolation=cv2.INTER_LINEAR)
            cv2.imwrite(os.path.join(typedir, img_name), img)

#通过裁剪完之后的统计csv里面找出某一类细胞，返回路径数组
def get_cell_lists_for_train(csvpath, cropdir, celltype):
    lists = []
    if not os.path.exists(csvpath):
        return False, lists
    df = pd.read_csv(csvpath)
    df1 = df[df['celltype'] == int(celltype)]
    for index, row in df1.iterrows():
        cellpath = os.path.join(cropdir, row['cell'])
        if os.path.exists(cellpath):
            lists.append(cellpath)
    return False, lists

#出入需要的细胞类型数组，传出一个字典
def get_datasets_for_train(csvpath, cellsdir, celltypes=[1, 7]):
    typetree = {}
    print(csvpath)
    if not os.path.exists(csvpath):
        return False, typetree
    for i in celltypes:
        ret, l = get_cell_lists_for_train(csvpath, cellsdir, i)
        key = str(i)
        typetree[key] = l
    return True, typetree

class cervical_autokeras():
    def __init__(self, jobid, jobdir):
        self.jid = jobid
        self.jobdir = jobdir
        self.scratchdir  = os.environ.get('SCRATCHDIR', 'scratch/')
        #训练、预测任务的顶层目录, 训练是按照时间随机生成的，预测是人为指定的
        self.ROOTPATH = os.path.join(self.scratchdir, self.jobdir) #每个任务的根目录
        #保存中间过程的目录
        self.TEMP_DIR = os.path.join(self.ROOTPATH, 'autokeras')
        #裁剪出来的细胞图存放的位置
        self.CELL_DIR = os.path.join(self.ROOTPATH, 'cells')
        self.TRAINJSON = os.path.join(self.ROOTPATH, 'train.json')
        self.MODJSON = os.path.join(self.ROOTPATH, 'mod.json')
        self.STATISTICS_DIR = os.path.join(self.ROOTPATH, 'statistics')
        self.CELL_CROP_CSV = os.path.join(self.STATISTICS_DIR, str(self.jid) + '_crop_cells.csv')
        self.CELL_CROP_DIR = os.path.join(self.CELL_DIR, 'crop')
        #Folder for storing training images
        self.TRAIN_IMG_DIR = os.path.join(self.ROOTPATH, 'train')
        self.RESIZE_TRAIN_IMG_DIR = os.path.join(self.ROOTPATH,'resize_train')
        #Folder for storing testing images
        self.TEST_IMG_DIR = os.path.join(self.ROOTPATH, 'test')
        self.RESIZE_TEST_IMG_DIR = os.path.join(self.ROOTPATH, 'resize_test')
        #Folder for storing predict images
        self.PREDICT_IMG_DIR = os.path.join(self.ROOTPATH, 'predict')
        self.RESIZE_PREDICT_IMG_DIR = os.path.join(self.ROOTPATH, 'resize_predict')
        self.PREDICT_ERROR_IMG_DIR = os.path.join(self.ROOTPATH, 'predict_error_data')
        #Path to generate csv file
        self.TRAIN_CSV_DIR = os.path.join(self.ROOTPATH, 'train_labels.csv')
        self.TEST_CSV_DIR = os.path.join(self.ROOTPATH, 'test_labels.csv')
        self.PREDICT_CSV_DIR = os.path.join(self.ROOTPATH, 'predict_labels.csv')
        #Path to generate model file
        self.MODEL_DIR = os.path.join(self.ROOTPATH, 'Modelak.h5')
        #If your memory is not enough, please turn down this value.(my computer memory 16GB)
        self.RESIZE = 128
        #Set the training time, this is half an hour
        self.TIME = 0.5*60*60
        self.percent = 0

        self.clean_fold()

    def clean_fold(self):
        dirs = [self.RESIZE_TRAIN_IMG_DIR, self.RESIZE_TEST_IMG_DIR, \
                self.RESIZE_PREDICT_IMG_DIR, self.PREDICT_ERROR_IMG_DIR, \
                self.TEMP_DIR]
        files = [self.TRAIN_CSV_DIR, self.TEST_CSV_DIR, self.PREDICT_CSV_DIR]
        for d in dirs:
            if os.path.exists(d):
                rmtree(d)
            os.makedirs(d)
        for f in files:
            if os.path.exists(f):
                os.remove(f)

    def train_autokeras(self):
        #Load images
        train_data, train_labels = load_image_dataset(csv_file_path=self.TRAIN_CSV_DIR, images_path=self.RESIZE_TRAIN_IMG_DIR)
        test_data, test_labels = load_image_dataset(csv_file_path=self.TEST_CSV_DIR, images_path=self.RESIZE_TEST_IMG_DIR)

        train_data = train_data.astype('float32') / 255
        test_data = test_data.astype('float32') / 255
        print("Train data shape:", train_data.shape)

        clf = ImageClassifier(verbose=True, path=self.TEMP_DIR)
        clf.fit(train_data, train_labels, time_limit=self.TIME)
        clf.final_fit(train_data, train_labels, test_data, test_labels, retrain=True)

        y = clf.evaluate(test_data, test_labels)
        print("Evaluate:", y)

        # clf.load_searcher().load_best_model().produce_keras_model().save(MODEL_DIR)
        # clf.export_keras_model(MODEL_DIR)
        clf.export_autokeras_model(self.MODEL_DIR)

    def predict_autokeras(self):
        #Load images
        test_data, test_labels = load_image_dataset(csv_file_path=self.PREDICT_CSV_DIR, images_path=self.RESIZE_PREDICT_IMG_DIR)
        test_data = test_data.astype('float32') / 255
        print("Test data shape:", test_data.shape)

        autokeras_model = pickle_from_file(self.MODEL_DIR)
        autokeras_score = autokeras_model.evaluate(test_data, test_labels)
        print(autokeras_score)

    def predict_autokeras2(self):
        autokeras_model = pickle_from_file(self.MODEL_DIR)

        #Load images
        for label in os.listdir(self.RESIZE_PREDICT_IMG_DIR):
            images = os.listdir(os.path.join(self.RESIZE_PREDICT_IMG_DIR, label))
            total = len(images)
            count_false = 0
            for index in range(0, total):
                img_path = os.path.join(self.RESIZE_PREDICT_IMG_DIR, label, images[index])
                if not os.path.exists(img_path):
                    continue

                img = load_img(img_path)
                x = img_to_array(img)
                x = x.astype('float32') / 255
                x = np.reshape(x, (1, self.RESIZE, self.RESIZE, 3))
                y = autokeras_model.predict(x)
                if str(label) != str(y[0]):
                    #print("%s %s result=%s" % (images[index], label, y[0]))
                    count_false = count_false + 1
                    error_image_dir = os.path.join(self.PREDICT_ERROR_IMG_DIR, label)
                    if not os.path.exists(error_image_dir):
                        os.makedirs(error_image_dir)
                    copyfile(img_path, os.path.join(error_image_dir, images[index]))
            print("%s 的个数/准确率：%d %f 出错的个数%d" % (label, total, (total - count_false) / total, count_false))
    def done(self, text):
        self.percent = 100
        post_job_status(self.jid, ds.TRAINNING_DONE.value, self.percent)
        print(text)
        return
    def processing(self, percent):
        self.percent = percent
        post_job_status(self.jid, ds.TRAINNING.value, self.percent)
        return
    def failed(self, text):
        post_job_status(self.jid, ds.TRAINNING_ERR.value, self.percent)
        print(text)
        return

if __name__ == "__main__":
    while 1:
        #向服务器请求任务，任务的状态必须是 READY4TRAIN
        if localdebug is not True and localdebug != "True":
            # datatype:  0未知1训练2预测
            jobid, status, dirname, jobtype = get_one_job(ds.READY4TRAIN.value, dt.TRAIN.value)
        else:
            jobid = 31
            status = ds.READY4TRAIN.value
            dirname = 'BVv1p1U6'

        print(jobid, status, dirname, jobtype)
        #检查得到的任务是不是想要的
        if status != ds.READY4TRAIN.value or dirname is None:
            time.sleep(5)
            continue

        ca = cervical_autokeras(jobid, dirname)
        #取出要训练的细胞类型
        jsonfile=load_json_file(ca.TRAINJSON)
        celltypes = jsonfile['types']
        ret, typetree = get_datasets_for_train(ca.CELL_CROP_CSV, ca.CELL_CROP_DIR, celltypes=celltypes)
        ca.processing(10)

        print ("Resize images...")
        resize_img2(typetree, ca.RESIZE_TRAIN_IMG_DIR, ca.RESIZE)
        resize_img2(typetree, ca.RESIZE_TEST_IMG_DIR, ca.RESIZE)
        print ("write csv...")
        write_csv(ca.RESIZE_TRAIN_IMG_DIR, ca.TRAIN_CSV_DIR)
        write_csv(ca.RESIZE_TEST_IMG_DIR, ca.TEST_CSV_DIR)
        print ("============Load...=================")
        ca.processing(20)
        ca.train_autokeras()

        #更新模型信息
        update_model_info_json(ca)
        ca.done("done!")
        #elif opt.task == 'predict':
        #    print ("Resize images...")
        #    resize_img(ca.PREDICT_IMG_DIR, ca.RESIZE_PREDICT_IMG_DIR, ca.RESIZE)
        #    print ("write csv...")
        #    write_csv(ca.RESIZE_PREDICT_IMG_DIR, ca.PREDICT_CSV_DIR)
        #    print ("============Load...=================")
        #    ca.predict_autokeras()

        #elif opt.task == 'predict2':
        #    print ("Resize images...")
        #    resize_img(ca.PREDICT_IMG_DIR, ca.RESIZE_PREDICT_IMG_DIR, ca.RESIZE)
        #    print ("write csv...")
        #    write_csv(ca.RESIZE_PREDICT_IMG_DIR, ca.PREDICT_CSV_DIR)
        #    print ("============Load...=================")
        #    ca.predict_autokeras2()

        del ca
        gc.collect()
        time.sleep(5)

        if localdebug is "True" or localdebug is True:
            break
