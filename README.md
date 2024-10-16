caching-proxy-server (https://roadmap.sh/projects/caching-server) is built using go lang and module used for server is github.com/gin-gonic/gin@latest and flag module for CLI arguments.
Currently Only supports open GET endpoints

Arguments
  - Port (Default: 80)
  - Origin (Required)
  - clear-cache


![image](https://github.com/user-attachments/assets/b6c2749a-bef5-4f1b-99f5-cd229fa27567)

![image](https://github.com/user-attachments/assets/94da279d-8dd2-49b0-8291-4765371afee0)

![image](https://github.com/user-attachments/assets/947a7f05-0ebc-413f-b304-c3353907ef88)

In first request X-Cache header will be missed

![image](https://github.com/user-attachments/assets/0d5302b1-0859-43fb-a379-7f8bb8c2b7dc)

And in subsequent request it will check for cached response by path in-memory dictionary (Map) 

![image](https://github.com/user-attachments/assets/7f8c2bc9-af19-41d0-b390-391dfd985be9)
