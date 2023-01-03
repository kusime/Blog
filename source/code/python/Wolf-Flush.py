import glob2
import asyncio
import os

from numpy import record
os.chdir(os.path.dirname(__file__))
iterators_of_file = []
dir_path_list = ['./_posts/狼与香辛料翻译/*']  # ['.\\*'] for windows
record_repeat = []


def make_loop(dir_path):
    sub1_round = glob2.glob(dir_path+'/*')  # 输出文件列表目录
    # +'\\*' for windows

    for sub1 in sub1_round:  # matched
        if sub1[-3:] == '.md':
            iterators_of_file.append(sub1)
        else:  # is dir
            dir_path_list.append(sub1)


for dir in dir_path_list:
    make_loop(dir)


# just prepare all of the path of .md file

rw_fileTaskList = []  # 创建文件阅读列表
fileobj = []  # creat return


async def retrun_file(filepath):  # 定义添加函数
    fileobj.append(open(filepath, mode='r', encoding='utf-8'))

loop = asyncio.get_event_loop()

for TaskPath in iterators_of_file:

    TASK = loop.create_task(retrun_file(TaskPath))
    rw_fileTaskList.append(TASK)


async def gorun(rw_fileTaskList):  # run task
    for task in rw_fileTaskList:
        await task
        #print("Running > ",task)

loop.run_until_complete(gorun(rw_fileTaskList))
# above is creat file object list


async def startfulush(fileobj_one):
    
    header = fileobj_one.readlines()[:10]# extra the header
    fileobj_one.seek(0) # reset pointer
    raw = fileobj_one.readlines()[10:] # pluck in lines exclude header
    fileobj_one.seek(0) # reset pointer
    
    raw_set = set(raw) # get unique but not shorted 
    # require except len < 5 比如说 ".." "是吗？"
    for key in list(raw_set):
        if len(key) < 14:
            raw_set.remove(key)
    # require except len < 10
    count_key = []
    for key in raw_set:
        count_key.append([key,raw.count(key)]) # can get key count
        # format [key,key_count]


    not_1_key = []
    for elemnet_count in count_key:
        if elemnet_count[1] != 1:
            not_1_key.append(elemnet_count)



    def reverse_flush(not_1_key,raw_input):
        for key in not_1_key:
            while key[1] != 1: # not finished
                raw_input.reverse()
                raw_input.pop(raw_input.index(key[0]))
                key[1]-=1
                raw_input.reverse()
                record_repeat.append(key[0])

    reverse_flush(not_1_key,raw)
    
    # raw is the api

    if '\n' not in raw:
        fileobj_one.seek(0)# reset pointer
        filtered_newline = raw
    else:
        fileobj_one.seek(0)# reset pointer
        filtered_newline = [x for x in raw if x != '\n']


    
    filtered_emp = [x for x in filtered_newline if x != '']
    filtered_space_line = [x for x in filtered_emp if x != ' \n']
    filtered_newline = filtered_space_line  # rewarp

    def rep(list, index, raw, new):
        list[index] = list[index].replace(raw, new)
        return list[index]
    
    
    filtered_russia = [rep(filtered_newline, filtered_newline.index(x), '俄罗斯', '罗伦斯')
                       for x in filtered_newline if '俄罗斯' in x]  # Russia in santance index
    print(fileobj_one.name.split('/')[-1],'filtered_russia',len(filtered_russia))
    filtered_holo = [rep(filtered_newline, filtered_newline.index(x), '赫罗', '赫萝')
                     for x in filtered_newline if '赫罗' in x]  
    print(fileobj_one.name.split('/')[-1],'filtered_holo',len(filtered_holo))
    recover_bolang = [rep(filtered_newline, filtered_newline.index(x), '\\~', '~')
                      for x in filtered_newline if '\\~' in x]  # Russia in santance index\                    
    print(fileobj_one.name.split('/')[-1],'recover_bolang',len(recover_bolang))
    replace_bolang = [rep(filtered_newline, filtered_newline.index(x), '~', '\~')
                      for x in filtered_newline if '~' in x]  # Russia in santance index\
    print(fileobj_one.name.split('/')[-1],'replace_bolang',len(replace_bolang))
    replace_keluo = [rep(filtered_newline, filtered_newline.index(x), '克罗', '赫萝')
                     for x in filtered_newline if '克罗' in x]  # Russia in santance index\
    print(fileobj_one.name.split('/')[-1],'replace_keluo',len(replace_keluo))
    replace_make = [rep(filtered_newline, filtered_newline.index(x), '马克', '马赫')
                    for x in filtered_newline if '马克' in x]  # Russia in santance index\
    print(fileobj_one.name.split('/')[-1],'replace_make',len(replace_make))
    replace_shangcheng = [rep(filtered_newline, filtered_newline.index(x), '商城', '商行')
                          for x in filtered_newline if '商城' in x]  # Russia in santance index\
    print(fileobj_one.name.split('/')[-1],'replace_shangcheng',len(replace_shangcheng))
    replace_mingren = [rep(filtered_newline, filtered_newline.index(x), '鸣人', '名人')
                       for x in filtered_newline if '鸣人' in x]  # Russia in santance index\
    print(fileobj_one.name.split('/')[-1],'replace_mingren',len(replace_mingren))
    replace_midiou = [rep(filtered_newline, filtered_newline.index(x), '米迪鸥', '米迪欧')
                      for x in filtered_newline if '米迪鸥' in x]
    print(fileobj_one.name.split('/')[-1],'replace_midiou',len(replace_midiou))
    replace_slash = [rep(filtered_newline, filtered_newline.index(x), '_', '---')
                     for x in filtered_newline if '_' in x]
    print(fileobj_one.name.split('/')[-1],'replace_slash',len(replace_slash))
    replace_keluo2 = [rep(filtered_newline, filtered_newline.index(x), '科罗', '赫萝')
                      for x in filtered_newline if '科罗' in x]
    print(fileobj_one.name.split('/')[-1],'replace_keluo2',len(replace_keluo2))
    replace_heluo = [rep(filtered_newline, filtered_newline.index(x), '贺罗', '赫萝')
                     for x in filtered_newline if '贺罗' in x]
    print(fileobj_one.name.split('/')[-1],'replace_heluo',len(replace_heluo))
    replace_title = [rep(filtered_newline, filtered_newline.index(x), 'Spice And Wolf Spice And Wolf Volume', 'Spice And Wolf Volume')
                     for x in filtered_newline if 'Spice And Wolf Spice And Wolf Volume' in x]
    print(fileobj_one.name.split('/')[-1],'replace_title',len(replace_title))
    replace_douhao = [rep(filtered_newline, filtered_newline.index(x), '，', ',')
                     for x in filtered_newline if '，' in x]
    print(fileobj_one.name.split('/')[-1],'replace_douhao',len(replace_douhao))
    replace_juhao = [rep(filtered_newline, filtered_newline.index(x), '。', '.')
                     for x in filtered_newline if '。' in x]
    print(fileobj_one.name.split('/')[-1],'replace_juhao',len(replace_juhao))
    
    
    name=fileobj_one.name
    fileobj_one.close() #cloue read file
    
    write_file = open(name,encoding='utf-8',mode='w') # reopen aovid error
    write_file.writelines(header+filtered_newline)
    write_file.close()
    print(name,"<<Flush Done")


TaskList = []

for file in fileobj:
    Task = loop.create_task(startfulush(file))
    TaskList.append(Task)

loop.run_until_complete(gorun(TaskList))

print('\n')
print("This Time Flushed Repeated Lines")
print('\n')
for line in record_repeat:
    print(line)

