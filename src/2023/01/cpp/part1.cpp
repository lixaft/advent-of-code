#include <iostream>
#include <string>

const char *digits = "1234567890";

int main() {
    int ret = 0;

    std::string line;
    while (getline(std::cin, line)) {
        std::string snum;
        snum += line[line.find_first_of(digits)];
        snum += line[line.find_last_of(digits)];

        ret += std::stoi(snum);
    }

    std::cout << ret << std::endl;

    return 0;
}
