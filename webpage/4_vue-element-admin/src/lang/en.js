export default {
  loginpage: {
    passwordRecovery: 'Password Recovery',
    register: 'Register',
    login: 'Welcome',
    inputemail: 'Please Input Your E-mail',
    password: 'Password',
    confirm: 'Confirm Password',
    code: 'Verification Code',
    send: 'Get Code',
    registing: 'Register And Login',
    gotologin: 'Login',
    gotoRecovery: 'Password Recovery',
    gotoRegister: 'Register',
    resetPassword: 'Reset Password',
    incorrectmail: 'Mailbox name not allowed (mailbox syntax incorrect)',
    incorrectpasswd: 'At least 3 letters or numbers',
    incorrectconfirm: 'The password and password (Reenter to confirm) do not match',
    enterCode: 'Please enter 6-digit email verification code'
  },
  web: {
    title: 'AI Medical'
  },
  navbar: {
    lang: 'English',
    logout: 'Logout'
  },
  route: {
    homepage: 'Dashboard',
    workSpace: 'WorkSpace',
    workSpaceOverview: 'WorkSpace Overview',
    workSpaceDetails: 'Details',
    report: 'Report',
    reportOverview: 'Report Overview',
    cell: 'Cells',
    cellOverview: 'Cells Overview',
    label: 'Label',
    system: 'System Settings',
    systemImg: 'Image',
    systemMail: 'E-Mail',
    systemErr: 'Error Log',
    permission: 'Permission',
    permissionUser: 'Users',
    permissionLog: 'Log',
    permissionInfo: 'User Info'
  },
  dashboard: {
    datasets: 'Datasets',
    images: 'Images',
    projects: 'Projects',
    users: 'Users',
    medicalID: 'Medical ID',
    medicalDesc: 'Description',
    medicalDataset: 'Dataset',
    medicalTS: 'Create Time',
    medicalETA: 'Status',
    reviewID: 'Review ID',
    reviewOwner: 'User',
    reviewUsername: 'Username',
    review1: 'Reviewed',
    review0: 'To Be Reviewed',
    datasetID: 'Dataset ID',
    datasetDesc: 'Description',
    datasetTS: 'Create Time',
    datasetETA: 'Status',
    modelID: 'Model ID',
    modelDesc: 'Description',
    modelType: 'Type',
    modelPrecision: 'Precision'
  },
  const: {
    init: 'Initialize',
    ready: 'Ready',
    start: 'Start',
    error: 'Error',
    done: 'Done',
    tobereview: 'To Be Review',
    reviewed: 'Reviewed',
    train: 'Train',
    predict: 'Predict',
    cellsType1: '1 Norm Cell N',
    cellsType2: '2 LSIL(Low-grade Squamous Intraepithelial Lesion) P',
    cellsType3: '3 HSIL(High-grade Squamous Intraepithelial Lesion) P',
    cellsType4: '4 HPV(Human papillomavirus) P',
    cellsType6: '6 SCC(Squamous Cell Carcinoma) P',
    cellsType7: '7 ASC-US(Atypical Squamous Cell Of Undetermined Significance) P',
    cellsType8: '8 ASC-H(Atypical Squamous Cell Cannot Exclude HSIL) P',
    cellsType9: '9 AGC(Atypical Glandular Cell) P',
    cellsType10: '10 AIS(Adenocarcinoma In Situ Of The Cervical Canal) P',
    cellsType11: '11 ADC(Adenocarcinoma) P',
    cellsType12: '12 T(Trichomonas) N',
    cellsType13: '13 M(Mould) N',
    cellsType14: '14 HSV(Herpes) N',
    cellsType15: '15 X1(Clue Cell) N',
    cellsType50: '50 Negative cells',
    cellsType51: '51 Positive cells',
    cellsType100: '100 Unknown',
    cellsType200: '200 Not Cell',
    medicalType50: 'Negative',
    medicalType51: 'Positive',
    medicalType100: 'Unknown',
    cellNorn: 'Normal',
    cellT: 'T(Trichomonas)',
    cellX1: 'X1(Clue Cell)',
    cellHSV: 'HSV(Herpes)',
    cellM: 'M(Mould)',
    cellASCUS: 'ASC-US(Atypical Squamous Cell Of Undetermined Significance)',
    cellASH: 'ASC-H(Atypical Squamous Cell Cannot Exclude HSIL)',
    cellLSIL: 'LSIL(Low-grade Squamous Intraepithelial Lesion)',
    cellHSIL: 'HSIL(High-grade Squamous Intraepithelial Lesion)',
    cellHPV: 'HPV',
    cellSCC: 'SCC(Squamous Cell Carcinoma)',
    cellAGC: 'AGC(Atypical Glandular Cell)',
    cellAIS: 'AIS(Adenocarcinoma In Situ Of The Cervical Canal)',
    cellADC: 'ADC(Adenocarcinoma)',
    cellNot: 'Not Cell',
    cellUnknown: 'Unknown',
    labelCellNorn: 'Normal Cell',
    labelCellGlandular: 'Glandular Cell',
    labelCellX1: 'Clue Cell',
    labelCellHSV: 'Herpes',
    labelCellM: 'Mould',
    labelCellPM: 'Atrophic Cellular Change',
    labelCellMetaplasia: 'Metaplasia',
    labelCellB: 'Biological',
    labelCellT: 'Trichomonas',
    labelCellASCUS: 'Atypical Squamous Cell Of Undetermined Significance',
    labelCellASCH: 'Atypical Squamous Cell Cannot Exclude HSIL',
    labelCellAGC: 'Atypical Glandular Cell',
    labelCellLSIL: 'Low-grade Squamous Intraepithelial Lesion',
    labelCellHSIL: 'High-grade Squamous Intraepithelial Lesion',
    labelCellHPV: 'Human papillomavirus',
    labelCellSCC: 'Squamous Cell Carcinoma',
    labelCellAIS: 'Adenocarcinoma In Situ Of The Cervical Canal',
    labelCellADC: 'Adenocarcinoma',
    labelCellNC: 'Not Cell',
    labelCellUnkown: 'Unkown',
    labelCellN: 'Negative',
    labelCellP: 'Positive'
  },
  workspace: {
    projects: 'Projects',
    datasets: 'Datasets',
    models: 'Models',
    projectsRefresh: 'Refresh',
    projectsNew: 'New Project',
    loading: 'Loading',
    projectDir: 'Directory',
    projectOwner: 'Creator',
    projectModelID: 'Model ID',
    projectModelType: 'Model Type',
    projectSize: 'Image Size',
    projectTime: 'Maximum Time Used For Training',
    projectType: 'Prediction Method',
    projectTypeLabel: 'With Label',
    projectTypeNoLabel: 'With No Label',
    projectPercent: 'Progress',
    projectETA: 'ETA(seconds)',
    projectStatus: 'Status',
    projectStartTime: 'Start Time',
    projectEndTime: 'End Time',
    projectID: 'Project ID',
    projectDesc: 'Description',
    projectDID: 'Dataset ID',
    projectCreator: 'Creator',
    projectCreatTime: 'Create Time',
    projectType2: 'Type',
    projectStatus2: 'Status',
    projectOP: 'Operation',
    projectOPLook: 'Details',
    projectNew: 'New Project',
    projectEnterDesc: 'Enter a description',
    projectStart: 'Start processing',
    projectNewType: 'Type',
    projectNewTypeTrain: 'Train',
    projectNewTypePredict: 'Predict',
    projectNewLabel: 'Label',
    projectNewNoLabel: 'No Label',
    projectNewTypeLabelled: 'Labelled',
    projectNewSize: 'Size(pixel)',
    projectNewTime: 'Maximum Time(seconds)',
    projectNewCache: 'Cache',
    projectNewNoCache: 'No Cache',
    projectNewColor: 'Color',
    projectNewGray: 'Gray',
    projectNewUser: 'Ordinary User',
    projectNewSelectDt: 'Select Dataset',
    projectNewSelect: 'Please Select',
    projectNewCreator: 'Creator',
    projectNewDir: 'Directory',
    projectNewStatus: 'Status',
    projectNewCropModel: 'Model Of Segmentation',
    projectNewUseCache: 'Use Cache',
    projectNewUseColor: 'Use Color Image',
    projectNewSize2: 'Size Of Segmentation',
    projectNewType2: 'Cell Type',
    projectNewEnterDesc: 'Enter A Description',
    projectSaveModel: 'Save Model',
    projectSelectModel: 'Select Model',
    projectSelectModel2: 'Select',
    projectModelPrecision: 'Precision',
    projectModelRecall: 'Recall',
    projectModelLoss: 'Loss',
    projectModelSize: 'Size',
    projectModelType2: 'Type',
    projectModelTrainNum: 'Number Of Training',
    projectModelClassNum: 'Classes',
    projectModelCellType: 'Cell Type',
    projectDataRefresh: 'Refresh',
    projectDataAdd: 'Add',
    projectDataUpload: 'Upload',
    projectDataCustom: 'Upload Custom Case',
    projectDataLoading: 'Loading',
    projectDataDir: 'Directory',
    projectDataCreator: 'Creator',
    projectDataModel: 'Model',
    projectDataCache: 'Cache',
    projectDataColor: 'Color',
    projectDataSize: 'Size',
    projectDataStartTime: 'Start Time',
    projectDataEndTime: 'End Time',
    projectDataID: 'ID',
    projectDataDescription: 'Description',
    projectDataCreator2: 'Creator',
    projectDataModel2: 'Model',
    projectDataCreateTime: 'Create Time',
    projectDataStatus: 'Status',
    projectDataOP: 'Operation',
    projectDataDetails: 'Details',
    projectDataUpload2: 'Upload',
    projectDataCustomUpload: 'Upload Custom Data',
    projectDataAdd2: 'New',
    projectDataPrevious: 'Previous',
    projectDataNext: 'Next',
    projectDataUpload3: 'Upload Images',
    projectDataSelect: 'Select Images',
    projectDataModel3: 'Select Model',
    projectDataReconfirm: 'Reconfirm',
    projectDataUploadTip: 'Drag & Drop Your Medical Case To Begin Uploading',
    projectDataBrowse: 'Select Folder',
    projectDataIncomplete: 'The Selected Case Is Incomplete',
    projectDataStartUpload: 'Uploading',
    dataSegmentationModel: 'Segmentation Model',
    dataSegmentationModel2: 'Select Model',
    dataCache: 'Use Cache',
    dataCacheYes: 'Yes',
    dataCacheNo: 'No',
    dataColor: 'Segmentation With Color',
    dataColorGray: 'Gray',
    dataColorColor: 'Color',
    dataSegmentationSize: 'Segmentation Size(pixel)',
    dataSegmentationType: 'Segmentation Type',
    dataSegmentationType0: '0-Cells were detected and cut out',
    dataSegmentationType1: '1-The cells were cut according to the labeled CSV',
    dataDescription: 'Enter A Description',
    dataStartToSegmentate: 'Start To Segmentate',
    dataTrainingSet: 'Training Set',
    dataBatch: 'Batch',
    dataCase: 'Case',
    dataDescription2: 'Description',
    dataProportion: 'n/p Proportion',
    dataModelParameters: 'Model And Parameters',
    dataModelID: 'Model ID',
    dataModel: 'Model',
    dataUseCache: 'Use Cache',
    dataColor2: 'Color',
    dataSize: 'Segmentation Size',
    dataUpdateTime: 'Update At',
    modelRefresh: 'Refresh',
    modelAdd: 'Add Model',
    modelLoading: 'Loading',
    modelID: 'Model ID',
    modelDesc: 'Description',
    modelType: 'Type',
    modelPrecision: 'Precision',
    modelCreatedAt: 'Created Time',
    modelRecall: 'Recall',
    modelLoss: 'Loss',
    modelNumber: 'Number Of Image',
    modelTypes: 'Types',
    modelUpload: 'Upload Model(Follow The Steps To Upload model)',
    modelUploadDesc: 'Please enter the model description',
    modelUploadDesc2: 'Please enter the description',
    modelUploadType: 'Please select model type',
    modelUploadType2: 'select model type',
    modelUploadUp: 'Select the model file to upload',
    modelUploadUp2: 'Select the model file',
    modelUploadUpcfg: 'Select the cfg file',
    modelUploadUp3: 'model file to be upload',
    modelUploadUpConfirm: 'Confirm information and upload',
    modelUploadUpConfirm2: 'Confirm && Upload',
    modelUploadUpFinished: 'Upload finished',
    modelUploadCancel: 'Cancel uploading',
    modelUploadRuleDesc: 'Please enter the model description',
    modelUploadRuleType: 'Please select model type',
    modelTypeSeg: 'Segmentation',
    modelTypeSegYOLOV4: 'Segmentation(YOLOV4)',
    modelTypeClassification: 'Classification',
    modelUploadOnlyH5: 'Only H5 / H5 files can be uploaded',
    dataType: 'Type',
    dataDir: 'Directory',
    dataBatch2: 'Batch',
    dataCase2: 'Case',
    dataNumberImage: 'Number Of Image',
    dataNumberN: 'fov-n Number',
    dataNumberP: 'fov-p Number',
    dataNumberLabels: 'The number of labels or cells',
    dataNumberNLabels: 'n Number',
    dataNumberPLabels: 'p Number',
    dataOriginalImage: 'Original Image',
    dataCells: 'Cells',
    dataSegmentationLog: 'Log',
    dataInProgress: 'In Progress',
    dataETA: 'Expected to take (sec.):',
    dataInputInformation: 'Input',
    dataOutputInformation: 'Output',
    dataDownloadAlert: 'Download in progress, please wait for download to complete! Please do not leave or refresh the page! Download size of compressed package(MB):',
    dataDownloadConfirm: 'Ok',
    dataDownloadZip: 'Cells.zip',
    dataDownloadSucc: 'Download successful',
    dataTrainingEvaluation: 'Training & Evaluation',
    dataImages: 'Images',
    dataPredict: 'Predict',
    dataDetails: 'Details',
    dataPredictDetails: 'Details',
    dataPredictNumber: 'Predict Number',
    dataPredictFalseNumber: 'False Number',
    dataPredictTrueNumber: 'True Number',
    dataPredictTrueCells: 'True Cells',
    dataPredictCells: 'Predict Cells',
    dataPredictFalseCells: 'False Cells',
    dataNotPredicted: 'Not Predicted',
    dataPredictOriginalImage: 'Original Image',
    dataPredictOverallProgress: 'Overall Progress',
    dataPredictStatus: 'Status',
    dataPredictTypes: 'Cell Types',
    dataPredictRefresh: 'Refresh',
    dataPredictCurrentProgress: 'Current Progress:',
    dataPredictFalse: 'False',
    dataPredictTrue: 'True',
    dataPredictTip: 'Tip',
    dataPredictInProgress: 'In Progress',
    dataPredictETA: 'Expected to take (sec.):',
    dataCustomUploaded: 'Uploaded',
    dataCustomNotUploaded: 'Not Uploaded',
    dataCustomFOVUploaded: 'Upload FOV',
    dataCustomUpload: 'Upload these images as a case',
    dataCustomNotFound: 'The picture file is not found or the format is not correct',
    dataCustomSizeNotUniform: 'The picture size is not uniform',
    dataCustomFormatUniform: 'The picture format is not uniform'
  },
  report: {
    doctorReview: 'Doctor Review',
    adminReview: 'Admin Review',
    loading: 'Loading',
    results: 'Results',
    directory: 'Directory',
    creator: 'Creator',
    modelID: 'Model ID',
    modelType: 'Model Type',
    Size: 'Size',
    MaxTime: 'Max Time',
    predictType: 'Predict Type',
    noLabel: 'No Label',
    labeled: 'Labeled',
    progress: 'Progress',
    ETA: 'Expected to take (sec.):',
    status: 'Status',
    startTime: 'Start Time',
    endTime: 'End Time',
    cellTypes: 'Cell Types',
    type: 'Type',
    number: 'Number',
    operation: 'Operation',
    download: 'Download',
    description: 'Description',
    datasetID: 'Dataset ID',
    creator1: 'Creator',
    type1: 'Type',
    createdAt: 'CreatedAt',
    ETA1: 'Expected to take (sec.):',
    operation1: 'Operation',
    review: 'Review'
  },
  label: {
    label: 'Label',
    loading: 'Loading',
    results: 'Results',
    directory: 'Directory',
    creator: 'Creator',
    modelID: 'Model ID',
    modelType: 'Model Type',
    Size: 'Size',
    MaxTime: 'Max Time',
    predictType: 'Predict Type',
    noLabel: 'No Label',
    labeled: 'Labeled',
    progress: 'Progress',
    ETA: 'Expected to take (sec.):',
    status: 'Status',
    startTime: 'Start Time',
    endTime: 'End Time',
    cellTypes: 'Cell Types',
    type: 'Type',
    number: 'Number',
    operation: 'Operation',
    download: 'Download',
    description: 'Description',
    datasetID: 'Dataset ID',
    creator1: 'Creator',
    type1: 'Type',
    createdAt: 'CreatedAt',
    ETA1: 'Expected to take (sec.):',
    operation1: 'Operation',
    review: 'Review',
    startDownloading: 'Start downloading',
    downloadFailedMsg: 'Download failed. Please refresh the page and try again',
    backPrevious: 'Back to previous page',
    total: 'Total',
    notReviewed: 'Not',
    reviewed: 'Reviewed',
    notReviewed2: 'Not Reviewed',
    cellReview: 'Cell Review',
    doctorReviewArea: 'Doctor review area',
    confirmation: 'Confirmation',
    negative: 'Negative',
    positive: 'Positive',
    other: 'Other',
    details: 'Details',
    operation2: 'Operation',
    label2: 'Label',
    exitLabelMode: 'Exit Label Mode',
    importSystemAnnotations: 'Import System Annotations',
    importedSystemAnnotations: 'Imported System Annotations',
    reviewSuccessful: 'Review Successful',
    topRight: 'Top Right',
    bottomRight: 'Bottom Right',
    bottomLeft: 'Bottom Left',
    topLeft: 'Top Left',
    buttonsRectangle: 'Draws a rectangular',
    actionsTitle: 'Cancel drawing rectangle',
    actionsText: 'Cancel',
    rectangleTooltipStart: 'Hold down the left button and drag the mouse to draw a rectangular',
    simpleshapeTooltipEnd: 'Release the left button to complete the drawing',
    buttonsEdit: 'Edit',
    buttonsEditDisabled: 'There are no annotations to Edit',
    buttonsRemove: 'Remove annotations',
    saveTitle: 'Save changes',
    saveText: 'Save',
    cancelTitle: 'Cancels the current annotations modification',
    cancelText: 'Cancel',
    clearAllTitle: 'Delete all labels, use with caution!',
    clearAllText: 'Delete all',
    removeDisabled: 'There are no annotations to delete',
    removeTooltipText: 'Click the annotation to delete it',
    editTooltipText: 'Drag the center point of the annotation to translate the annotation, and drag the four corners to change the annotation shape',
    editTooltipSubtext: '"Cancel / save" in the upper right corner can discard / save the changes'
  },
  system: {
    errorUser: 'User',
    errorOperationTime: 'Operation Time',
    errorVersion: 'Version',
    errorDomain: 'Domain',
    errorLog: 'Log',
    errorOperation: 'Operation',
    errorUserID: 'User ID',
    errorUserName: 'User Name',
    errorPhone: 'Phone',
    errorReqSize: 'Request Size',
    errorState: 'State',
    errorTime: 'Time(ms)',
    errorOperationTime2: 'Operation Time',
    errorRequestMethod: 'Request Method',
    errorPath: 'Path',
    errorFrom: 'From',
    errorCountry: 'Country',
    errorProvince: 'Province',
    errorCity: 'City',
    errorOperators: 'Operators',
    errorOS: 'OS',
    errorBrowser: 'Browser',
    errorType: 'Type',
    errorEMail: 'E-Mail',
    errorDetails: 'Details',
    imgPictureAntiTheftChain: 'Whether to enable picture anti-theft chain',
    imgEnabel: 'Enabel',
    imgDisabel: 'Disabel',
    imgCacheTime: 'image cache time',
    imgSelect: 'Please select',
    img401: 'Picture 401 settings',
    img404: 'Picture 404 settings',
    imgUpload: 'Upload',
    imgWhiteList: 'White list',
    imgEnterContent: 'Please enter the content',
    imgAdd: 'Add',
    imgSave: 'Save',
    imgReset: 'Reset',
    usrUploadAvatar: 'Upload Avatar',
    usrNickname: 'Nickname',
    usrEmail: 'Email',
    usrGender: 'Gender',
    usrGenderSelect: 'Please select your gender',
    usrMale: 'Male',
    usrFemale: 'Female',
    usrSecrecy: 'Secrecy',
    usrPhone: 'Phone',
    usrCreateAt: 'CreateAt',
    usrIntroduce: 'Introduce',
    usrSave: 'Save',
    usrReset: 'Reset',
    usrEnterNickname: 'Please enter your nickname',
    usrNicknameTip: 'The length is between 2 and 40 characters',
    usrEnterPhone: 'Please input your mobile phone number (optional)',
    usrPhoneTip: 'The length is 0 to 11 characters, which can be left blank',
    usrEnterIntroduce: 'Please enter self introduction',
    usrIntroduceTip: 'The length is 0 to 512 characters',
    usrUpdated: 'User information has been updated',
    usrFillAgain: 'Please fill in again',
    usrID: 'usr ID',
    usrAvatar: 'Avatar',
    usrName: 'Name',
    usrType: 'Type',
    usrCreatedAt: 'Created At',
    usrUpdatedAt: 'Updated At',
    usrOperation: 'Operation',
    usrPermission: 'Permission',
    usrDetails: 'Details',
    usrJustNow: 'Just Now',
    usrMinuteAgo: 'Minutes Ago',
    usrHoursAgo: 'Hours Ago',
    usr1DayAgo: '1 Day Ago',
    usrHW: 'HardWare',
    usrRegionID: 'Region ID',
    usrBrowserEngine: 'Browser Engine'
  }
}
