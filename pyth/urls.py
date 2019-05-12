import requests, re, sys

def all_urls(main_link):
    urls_domain=[main_link]; urls_outside=[]
    index=0
    while index<len(urls_domain):
        # put the coode to execute in every page here
        req = requests.get(urls_domain[index].split(' ')[0])
        print(urls_domain[index])
        print(req)
        urls_domain[index]=urls_domain[index]+' '+str(req)
        for link in re.findall('<a href=[\'"]?([^\'" >]+)', req.text):
            if link[0]=='/':
                url = main_link+link
                if url not in urls_domain:
                    urls_domain+=[url]
            elif ('http'==link[0:4]) and (link not in urls_domain):
                if main_link in link:
                    urls_domain+=[link]
                else:
                    urls_outside+=[link]
        index+=1
    return urls_domain, urls_outside;


main_url = sys.argv[1] if sys.argv[1][-1]!='/' else sys.argv[1][0:-1]
urls_domain, urls_outside = all_urls(main_url)

with open("domain_urls.txt", "w+") as f1:
    f1.write('\n'.join(urls_domain))

with open("outside_urls.txt", "w+") as f2:
    f2.write('\n'.join(urls_outside))
