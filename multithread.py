import threading
import time
import subprocess

myScale = 8

cmd = "bash ./invoke_storeAsset_clientSDK.sh"

class ScaleClient:

    def multiclient(self):
        returned_value = subprocess.call(cmd, shell=True)
        print('returned value:', returned_value)

    def __init__(self):
        t = threading.Thread(target=self.multiclient)
        t.start()

for i in range(myScale):
    ScaleClient()



#ScaleClient()
#ScaleClient()
