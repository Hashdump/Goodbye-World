import threading
import os

def d2c(l):
    threading.current_thread()
    os.getpid()
    print(chr(l), end = '')

if __name__ == "__main__":
    os.getpid()
    threading.main_thread()

    d = [72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100, 33]
    ts = []
    
    for i in range(len(d)):
        t = threading.Thread(target = d2c, args = (d[i],))
        ts.append(t)
        t.start()
        
    for j in range(len(d)):
        ts[j].join()
