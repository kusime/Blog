import glob2
import asyncio
import os
os.chdir(os.path.dirname(__file__))

iterators_of_file = []

dir_path_list = ['./*']##['.\\*'] for windows

def make_loop(dir_path):
    sub1_round = glob2.glob(dir_path+'/*')  # 输出文件列表目录
                                            # +'\\*' for windows

    for sub1 in sub1_round: #matched
        if sub1[-3:] == '.md' :
            iterators_of_file.append(sub1)
        else:# is dir
            dir_path_list.append(sub1)


for dir in dir_path_list:
    make_loop(dir)


# just prepare all of the path of .md file

rw_fileTaskList = []  # 创建文件阅读列表
fileobj = []  # creat return


async def retrun_file(filepath):  # 定义添加函数
    fileobj.append(open(filepath, mode='r+', encoding='utf-8'))


loop = asyncio.get_event_loop()

for TaskPath in iterators_of_file:

    TASK = loop.create_task(retrun_file(TaskPath))
    rw_fileTaskList.append(TASK)


async def gorun(rw_fileTaskList):
    for task in rw_fileTaskList:
        await task
        #print("Running > ",task)

loop.run_until_complete(gorun(rw_fileTaskList))


async def startfulush(fileobj_one):
    lines = fileobj_one.readlines()
    date_index = 0

    for line in lines[:20]:  # loop for upper 20 lines
        check = line.split(': ')[0]
        if  check == 'updated':# check if have been inseted ?
            return 1 # quit 


    for line in lines[:20]:  # loop for upper 20 lines
        check = line.split(': ')[0]
        if check == 'date':
            date_index = lines.index(line)
            break
        else:
            pass

    normal_date = lines[date_index].split(': ')
    print(normal_date)
    try:
        formal_check = normal_date[0]
        times = normal_date[1]
    except IndexError:
        print(fileobj_one)
        raise IndexError

    fileobj_one.seek(0)  # reset pointer
    newline = 'updated: ' + times
    print(newline)
    lines.insert(date_index+1, newline)
    fileobj_one.writelines(lines)
    fileobj_one.close()
    print(fileobj_one.name, "  Finished Insert")

TaskList = []

for file in fileobj:
    Task = loop.create_task(startfulush(file))
    TaskList.append(Task)

loop.run_until_complete(gorun(TaskList))
