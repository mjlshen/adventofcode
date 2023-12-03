file = open('input.txt','r').readlines()

for i,line in enumerate(file):
  file[i] = line.replace(': ', ' = ')

done = False
while not done:
  done = True
  for line in file:
    try:
      exec(line)
    except:
      done = False

print(root)
