// gmtime 获取格林威治时间。         //例如 time_t = 0 代表的时间是1970.1.1，格林威治0:0:0    
// localtime/time 获取当前时区时间。 //例如 time_t = 0 东8区则是1970.1.1 8:00:00

//如果没有设置时区gmtime与localtime/time时间就会不一致
//如果设置时区timezone-0，则gmtime与localtime/time就会一致了,都是获取的格林威治时间

//linux版本使用tzset、windows版本使用_tzset设置时区
//_putenv_s("TZ", "GMT-0");
//_tzset();

time_t curtime = 0;//0x000000004F9A54E5
//time(&curtime);

char* tt =   ctime(&curtime);//当地时区时间的字符串

tm* t = gmtime(&curtime);// gm为格林威治时间，如果不设置时区，则和ctime不一致

//0x000000004F9A54E5
//Fri Apr 27 10:12:21 CEST 2012  
