import time
import threading

counter = 0


def check_counter():
    global counter
    counter += 1
    print "current value:",counter
    print "Name:",threading.currentThread().getName()
    print time.ctime()



if __name__ == '__main__':
    for i in range(5):
        t = threading.Thread(target=check_counter, name="T"+str(i))
        t.start()
        
