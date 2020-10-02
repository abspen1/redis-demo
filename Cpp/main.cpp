#include <iostream>
#include <string>
#include <sw/redis++/redis++.h>

using namespace std;

using namespace sw::redis;



int main () {
    string password = getenv ("REDIS_PASS");
    if (password != "") {
        cout << password;
    }
    try
    {
        ConnectionOptions connection_options;
        connection_options.host = "10.10.10.1"; // Required.
        connection_options.port = 6379;        // Optional. The default port is 6379.
        connection_options.password = password;  // Optional. No password by default.
        connection_options.db = 1;             // Optional. Use the 0th database by default.
        Redis redis(connection_options);

        redis.set("key", "value");
        auto val = redis.get("key");
        if (val)
        {
            // dereference val to get the value of string type.
            std::cout << *val << std::endl;
        } // else key doesn't exist

        // Write elements in STL container to Redis
        redis.rpush("list", {"a", "b", "c"});

        // std::vector<std::string> vec = {"d", "e", "f"};
        // redis.rpush("list", vec.begin(), vec.end());

        // Write elements in Redis list to STL container
        // std::vector<std::string> res;
        // redis.lrange("list", 0, -1, std::back_inserter(res));
    }
    catch (const Error &e)
    {
        // Error handling
    }
}