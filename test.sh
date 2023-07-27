curl --location --request GET 'http://0.0.0.0:8080/community/page/get/2'
curl -H “Content-Type： application/json” -X POST -d '{"content":"hope success"}'“http://172.0.0.1:8080/community/page/post/”
curl --data "content=1111" http://172.0.0.1:8080/community/page/post
