import os 

with open(".env", "a") as f:
	f.write("\nAPI_SECRET=" + str(os.urandom(32).hex()))