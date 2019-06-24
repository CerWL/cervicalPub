import os
from sklearn.model_selection import KFold
import numpy as np
from segmentation_functions import *
from gan_model import *
import multiprocessing
import time
import pandas as pd

def cell_segmentation(positive_images_root, negative_images_root, positive_npy_root, 
                      negative_npy_root, ref_path, intensity, multi_core):
    
    positive_images_path = [positive_images_root + n for n in os.listdir(positive_images_root)]
    negative_images_path = [negative_images_root + n for n in os.listdir(negative_images_root)]
    
    if 1- os.path.exists(positive_npy_root+str(intensity)+'/'):
        os.makedirs(positive_npy_root+str(intensity)+'/')
    if 1- os.path.exists(negative_npy_root+str(intensity)+'/'):
        os.makedirs(negative_npy_root+str(intensity)+'/')

    if (multi_core == True):
        jobs = []
        for index, i in enumerate(positive_images_path):
            p = multiprocessing.Process(target=cell_segment, args=(i, positive_npy_root+str(intensity)+'/',
                                                                  ref_path, intensity))
            p.start()
            jobs.append(p)
            if (index+1)%7 == 0:
                p.join()
                jobs = []

        for job in jobs:
            p.join()

        jobs = []
        for index, i in enumerate(negative_images_path):
            p = multiprocessing.Process(target=cell_segment, args=(i, negative_npy_root+str(intensity)+'/',
                                                                  ref_path, intensity))
            p.start()
            jobs.append(p)
            if (index+1)%7 == 0:
                p.join()
                jobs = []
            
        for job in jobs:
            p.join()
    else:
        for index, i in enumerate(positive_images_path):
            cell_segment(i, positive_npy_root+str(intensity)+'/',ref_path,intensity)
        for index, i in enumerate(negative_images_path):
            cell_segment(i, negative_npy_root+str(intensity)+'/',ref_path,intensity)

def split_dataset(path, fold=4, random_seed=42):
    np.random.seed(random_seed)
    kf = KFold(n_splits=fold, shuffle=True)
    kf.get_n_splits(path)
    train_list, test_list, train_label_list, test_label_list  = [], [], [], []
    for train_index, test_index in kf.split(path):
        train_list.append([path[n] for n in train_index])
        test_list.append([path[n] for n in test_index])
    return train_list, test_list

def cell_representation(X_train_path, X_test_path, y_train_path, y_test_path, experiment_root, 
                        n_epoch=50, batchsize=16, rand=32, dis=1, dis_category=5, 
                        ld = 1e-4, lg = 1e-4, lq = 1e-4, save_model_steps=100):

    X_train = np.load(X_train_path)
    X_test = np.load(X_test_path)
    y_train = np.load(y_train_path)
    y_test = np.load(y_test_path)

    cell_train_set = np.concatenate([X_train, X_test])
    cell_test_set = cell_train_set
    cell_test_label = np.concatenate([y_train, y_test])

    netD, netG, netD_D, netD_Q = create_model(rand=rand, dis_category=dis_category)
    train_representation(cell_train_set, cell_test_set, cell_test_label, netD, netG, netD_D, netD_Q,
                         experiment_root, n_epoch=n_epoch, batchsize=batchsize, rand=rand, dis=dis, 
                         dis_category=dis_category, ld=ld, lg=lg, lq=lq, save_model_steps=save_model_steps)

def cell_representation_eval(X_train_path, X_test_path, y_train_path, y_test_path, experiment_root, 
                        n_epoch=50, batchsize=16, rand=32, dis=1, dis_category=5, 
                        ld = 1e-4, lg = 1e-4, lq = 1e-4, save_model_steps=100):

    X_train = np.load(X_train_path)
    X_test = np.load(X_test_path)
    y_train = np.load(y_train_path)
    y_test = np.load(y_test_path)

    cell_train_set = np.concatenate([X_train, X_test])
    cell_test_set = cell_train_set
    cell_test_label = np.concatenate([y_train, y_test])

    netD, netG, netD_D, netD_Q = create_model(rand=rand, dis_category=dis_category)
    eval_representation(cell_train_set, cell_test_set, cell_test_label, netD, netG, netD_D, netD_Q,
                         experiment_root, n_epoch=n_epoch, batchsize=batchsize, rand=rand, dis=dis, 
                         dis_category=dis_category, ld=ld, lg=lg, lq=lq, save_model_steps=save_model_steps)

    
def image_classification(positive_images_root, negative_images_root, positive_npy_root,negative_npy_root, 
                      ref_path, intensity, X_train_path, X_test_path, y_train_path, y_test_path, experiment_root, multi_core = True,
                      fold = 4, random_seed=42, choosing_fold = 1, n_epoch=10000, batchsize=32, rand=64, dis=1, 
                         dis_category=5, ld = 1e-4, lg = 1e-4, lq = 1e-4, save_model_steps = 100):

    #cell_segmentation(positive_images_root, negative_images_root, positive_npy_root, 
    #                  negative_npy_root, ref_path, intensity, multi_core)
    
    positive_npy_path = [positive_npy_root +str(intensity)+'/' + n[:-3] + 'npy' for n in os.listdir(positive_images_root)]
    negative_npy_path =[negative_npy_root +str(intensity)+'/' + n[:-3] + 'npy' for n in os.listdir(negative_images_root)]

    positive_train_list, positive_test_list = split_dataset(positive_npy_path, fold, random_seed)
    negative_train_list, negative_test_list = split_dataset(negative_npy_path, fold, random_seed)

    positive_train_npy = [np.load(n) for n in positive_train_list[choosing_fold]]
    positive_test_npy = [np.load(n) for n in positive_test_list[choosing_fold]]
    negative_train_npy = [np.load(n) for n in negative_train_list[choosing_fold]]
    negative_test_npy = [np.load(n) for n in negative_test_list[choosing_fold]]
    
    #print(np.array(positive_train_npy).shape, np.array(negative_train_npy).shape)
    print(np.concatenate(positive_train_npy).shape, np.concatenate(negative_train_npy).shape)
    cell_train_set = np.concatenate([np.concatenate(positive_train_npy), np.concatenate(negative_train_npy)])

    X_train = np.load(X_train_path)
    X_test = np.load(X_test_path)
    y_train = np.load(y_train_path)
    y_test = np.load(y_test_path)
    cell_test_set = np.concatenate([X_train, X_test])
    cell_test_label = np.concatenate([y_train, y_test])
    cell_train_set = np.concatenate([cell_train_set, rotation(cell_test_set)])

    netD, netG, netD_D, netD_Q = create_model(rand=rand, dis_category=dis_category)
    netD, netG, netD_D, netD_Q =  train(cell_train_set, cell_test_set, cell_test_label, 
                   positive_train_npy, positive_test_npy,negative_train_npy, negative_test_npy,
                   netD, netG, netD_D, netD_Q, experiment_root, 
                   n_epoch=n_epoch, batchsize=batchsize, rand=rand, dis=1, dis_category=dis_category, 
                   ld = ld, lg = lg, lq = lq, save_model_steps=save_model_steps)
                   
                   

      
                   
def image_classification_segment(positive_images_root, negative_images_root, positive_npy_root,negative_npy_root, 
                        positive_test_images_root, negative_test_images_root, positive_test_npy_root,negative_test_npy_root,
                      ref_path, intensity, X_train_path, X_test_path, y_train_path, y_test_path, experiment_root, multi_core = True,
                      fold = 4, random_seed=42, choosing_fold = 1, n_epoch=10000, batchsize=32, rand=64, dis=1, 
                         dis_category=5, ld = 1e-4, lg = 1e-4, lq = 1e-4, save_model_steps = 100):

    cell_segmentation(positive_test_images_root, negative_test_images_root, positive_test_npy_root, 
                      negative_test_npy_root, ref_path, intensity, multi_core)
                      
                      
def image_classification_train(positive_images_root, negative_images_root, positive_npy_root,negative_npy_root, 
                      ref_path, intensity, X_train_path, X_test_path, y_train_path, y_test_path, experiment_root, multi_core = True,
                      fold = 4, random_seed=42, choosing_fold = 1, n_epoch=10000, batchsize=32, rand=64, dis=1, 
                         dis_category=5, ld = 1e-4, lg = 1e-4, lq = 1e-4, save_model_steps = 100):

    #cell_segmentation(positive_test_images_root, negative_test_images_root, positive_test_npy_root, 
    #                  negative_test_npy_root, ref_path, intensity, multi_core)
    
    positive_npy_path = [positive_npy_root +str(intensity)+'/' + n[:-3] + 'npy' for n in os.listdir(positive_images_root)]
    negative_npy_path =[negative_npy_root +str(intensity)+'/' + n[:-3] + 'npy' for n in os.listdir(negative_images_root)]

    positive_train_list, positive_test_list = split_dataset(positive_npy_path, fold, random_seed)
    negative_train_list, negative_test_list = split_dataset(negative_npy_path, fold, random_seed)

    positive_train_npy = [np.load(n) for n in positive_train_list[choosing_fold]]
    negative_train_npy = [np.load(n) for n in negative_train_list[choosing_fold]]
    

    positive_test_npy = [np.load(n) for n in positive_test_list[choosing_fold]]
    negative_test_npy = [np.load(n) for n in negative_test_list[choosing_fold]]
    
    
    
    #print(np.array(positive_train_npy).shape, np.array(negative_train_npy).shape)
    print(np.concatenate(positive_train_npy).shape, np.concatenate(negative_train_npy).shape)
    cell_train_set = np.concatenate([np.concatenate(positive_train_npy), np.concatenate(negative_train_npy)])

    netD, netG, netD_D, netD_Q = create_model(rand=rand, dis_category=dis_category)
    netD, netG, netD_D, netD_Q, eval_vect =  train_predict(positive_train_npy, positive_test_npy,negative_train_npy, negative_test_npy,
                   netD, netG, netD_D, netD_Q, experiment_root, 
                   n_epoch=n_epoch, batchsize=batchsize, rand=rand, dis=1, dis_category=dis_category, 
                   ld = ld, lg = lg, lq = lq, save_model_steps=save_model_steps)
    
    
    print("Generate result:")
    train_feat, train_predict_label, test_feat, test_predict_label = eval_vect               
    train_files = positive_train_list[choosing_fold] + negative_train_list[choosing_fold]
    train_true_label = len(positive_train_list[choosing_fold])*[1] + len(negative_train_list[choosing_fold])*[0]
    print(len(train_feat), len(train_files), len(train_predict_label), len(train_true_label))
    train_df = pd.DataFrame({'ImageName':  [os.path.basename(file).split('.')[0] for file in train_files],
                             'FeatureVector': [feat for feat in train_feat],
                             'PredictLabel': [pl for pl in train_predict_label],
                             'TrueLabel': train_true_label,
                             'Dataset': ['Train'] * len(train_files)
                            })
    
               
    test_files = positive_test_list[choosing_fold] + negative_test_list[choosing_fold]
    test_true_label = len(positive_test_list[choosing_fold])*[1] + len(negative_test_list[choosing_fold])*[0]
    test_df = pd.DataFrame({'ImageName':  [os.path.basename(file).split('.')[0] for file in test_files],
                             'FeatureVector': [feat for feat in test_feat],
                             'PredictLabel': [pl for pl in test_predict_label],
                             'TrueLabel': test_true_label,
                             'Dataset': ['Test'] * len(test_files)
                             })
    res_df = train_df.append(test_df, ignore_index=True)
    res_df.to_csv("./svm_result.csv")
    
    
def image_classification_predict(positive_test_images_root, negative_test_images_root, positive_test_npy_root,negative_test_npy_root,
                      ref_path, intensity, experiment_root, rand=64, dis_category=5):


    positive_test_npy_path = [positive_test_npy_root +str(intensity)+'/' + n[:-3] + 'npy' for n in os.listdir(positive_test_images_root)]
    negative_test_npy_path =[negative_test_npy_root +str(intensity)+'/' + n[:-3] + 'npy' for n in os.listdir(negative_test_images_root)]

    positive_test_list = positive_test_npy_path
    negative_test_list = negative_test_npy_path

    positive_test_npy = [np.load(n) for n in positive_test_list]
    negative_test_npy = [np.load(n) for n in negative_test_list]
    

    netD, netG, netD_D, netD_Q = create_model(rand=rand, dis_category=dis_category)
    netD, netG, netD_D, netD_Q, predict_label, feature_dics =  predict(positive_test_npy,  negative_test_npy,
          netD, netG, netD_D, netD_Q, experiment_root,dis_category=dis_category)
    
    
    print("Positive image predict result:")
    positive_files = os.listdir(positive_test_images_root)
    positive_cnt = len(positive_files)
    
    if positive_cnt > 100:
        print('The positive images number is greater than 100. Please check the ./positive_res.csv')
    else:
        for image_file, pred_l in zip(positive_files, predict_label[:positive_cnt]):
            basename = image_file.split('.')[0]
            print("Image Name: {}; Predict Label: {}".format(basename, 'P' if pred_l else 'N'))
    
    res_df = pd.DataFrame({'Image Name': [file.split('.')[0] for file in positive_files], 
                           'Predict Label': ['P' if pl else 'N' for pl in predict_label[:positive_cnt] ],
                           'Feature Vector':[feat for feat in feature_dics[:positive_cnt]]
    })
    res_df.to_csv('./positive_res.csv')
    
    print("Negative image predict result:")
    negative_files = os.listdir(negative_test_images_root)
    negative_cnt = len(negative_files)
    
    if negative_cnt > 100:
        print('The negative images number is greater than 100. Please check the ./negative_res.csv')
    else:
        for image_file, pred_l in zip(negative_files, predict_label[-negative_cnt:]):
            basename = image_file.split('.')[0]
            print("Image Name: {}; Predict Label: {}".format(basename, 'P' if pred_l else 'N'))
            
    
    res_df = pd.DataFrame({'Image Name': [file.split('.')[0] for file in negative_files], 
                           'Predict Label': ['P' if pl else 'N' for pl in predict_label[-negative_cnt:] ],
                           'Feature Vector':[feat for feat in feature_dics[-negative_cnt:]]
    })
    res_df.to_csv('./negative_res.csv')
            
            
        
             