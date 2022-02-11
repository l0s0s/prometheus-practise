from random import randint, randrange
from time import sleep
from requests import get


urls = [
    'http://localhost:8080/api/v1/mock/ok',
    'http://localhost:8080/api/v1/mock/unauthorised',
    'http://localhost:8080/api/v1/mock/notfound'
]


def make_requests():
    while True:
        url = urls[randrange(len(urls))]

        get(url)
        print("request to {0}".format(url))

        sleep(randint(1, 5))


if __name__ == '__main__':
    make_requests()