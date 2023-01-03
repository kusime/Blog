import os,sys
import glob2


gened = [] # format = (date,filtered)
modify_gened_date = []


os.chdir(os.path.dirname(sys.argv[0]))

for line in glob2.iglob('./gallery/paint/*.*g'):
    line_time = os.path.getmtime(line)
    line = '![](/gallery/paint/'+line.split('/')[3]+')\n' # line.split('\\')[1]+')for windows 
                                    #'/' for linux
    gened.append((line_time,line))
    modify_gened_date.append(line_time)

gened_dict = dict(gened)
modify_gened_date.sort(reverse=False)# this pragram could change because this just effect the new picture inserting
#because the check existent function is nonsequenced check

gened = [] # reset gened,purpose is keep the variable is not change ,we just modity it's sequence
for date in modify_gened_date:
    gened.append(gened_dict[date])


def get_insert_index(lines:list,target:str):
    for line in lines[:20]:
        if line == target:
            return lines.index(line)


with open('./paint/index.md',mode='r+',encoding='utf-8') as fileobj:
    lines=fileobj.readlines()
    #print(lines)
    target = '{% gp -4 %}\n'
    insert_index = get_insert_index(lines,target)

    headnote = lines[:insert_index+1]

    previous_pic= lines[insert_index+1:-1]#get raw
    endnote =lines[-1]
    
    for line in gened:
        if line in previous_pic:
            pass
        else:
            print("New Picture Detected .. ", line)
            previous_pic.insert(0,line)
    previous_pic.append(endnote)
    
    after = headnote + previous_pic

    fileobj.seek(0)
    fileobj.writelines(after)
