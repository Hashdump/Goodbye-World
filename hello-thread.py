import threading
import os

def letter(l):
    threading.current_thread()
    os.getpid()
    print(chr(l), end = '')

if __name__ == "__main__":
 
    os.getpid()
    threading.main_thread()
 
    arr = [72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100, 33]
    threads = []
    
    for x in range(len(arr)):
        t = threading.Thread(target = letter, args=(arr[x],))
        threads.append(t)
        t.start()
        
    for z in range(len(arr)):
        threads[z].join()
