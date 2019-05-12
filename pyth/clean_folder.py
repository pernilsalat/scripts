import os, sys, errno

if(len(sys.argv)!=2):
    print('Usage: python3 '+sys.argv[0]+' /absolute/path/to/folder')
    sys.exit()

base_dir= sys.argv[1] if sys.argv[1][-1]=='/' else sys.argv[1]+'/'

def move(base_dir, gruped_files):
    for dir_name in gruped_files.keys():
        try:
            if gruped_files[dir_name]!=[]:
                os.makedirs(base_dir+dir_name)
        except OSError as e:
            if e.errno != errno.EEXIST:
                raise
        for file_name in gruped_files[dir_name]:
            os.rename(base_dir+file_name, base_dir+dir_name+'/'+file_name)

def organize_files(base_dir):
    files=[]; dir=[]
    gruped_files = {}
    formats = {
        'compresed_files' : ['.7z', '.tgz', '.deb', '.zip', '.apk', '.rar', '.tar'],
        'doc'             : ['.txt', '.rtf', '.pdf', '.doc', '.xls', '.xlm', '.ppt'],
        'programacio'     : ['.py', '.rb', '.java', '.js', '.c', '.html', '.go', '.json', '.sh', '.form'],
        'images'          : ['.jpeg', '.jpg', '.tiff', '.gif', '.png'],
    }
    for i in formats.keys():
        gruped_files[i] = []

    for name in os.listdir(base_dir):
        if os.path.isdir(base_dir+'/'+name):
            dir+=[name]
        else:
            for fmts in formats.keys():
                if any(fmt in name for fmt in formats[fmts]):
                    gruped_files[fmts].append(name)
                    break
            else:
                fmt = name.split('.')[-1]
                if fmt not in gruped_files.keys():
                    gruped_files[fmt] = []
                gruped_files[fmt].append(name)
            files+=[name]

    move(base_dir, gruped_files)
    return dir, files, gruped_files

dir, files, gruped_files = organize_files(base_dir)
print(dir, files, gruped_files, sep='\n\n')
