import requests
import regex as re

r = requests.get('http://www.pythonchallenge.com/pc/def/equality.html')

pattern = r'<!--(.*?)-->'
data = re.findall(pattern, r.text, re.S|re.M)[0]

for i in range(4, len(data)-4):
    if data[i-4].islower() and data[i-3].isupper() and data[i-2].isupper() and data[i-1].isupper() and data[i].islower() and data[i+1].isupper() and data[i+2].isupper() and data[i+3].isupper() and data[i+4].islower():
        print(data[i-3], data[i-2], data[i-1], data[i], data[i+1], data[i+2], data[i+3])