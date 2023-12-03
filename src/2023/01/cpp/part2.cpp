#include <iostream>
#include <string>

const char *digits = "0123456789";

void replace_all(
    std::string &line, const std::string &from, const std::string &to) {
    int i;
    while ((i = line.find(from)) != std::string::npos) {
        line.replace(i, from.length(), to);
    }
}

void replace_by_numbers(std::string &line) {
    replace_all(line, "zero", "z0o");
    replace_all(line, "one", "o1e");
    replace_all(line, "two", "t2o");
    replace_all(line, "three", "t3e");
    replace_all(line, "four", "f4r");
    replace_all(line, "five", "f5e");
    replace_all(line, "six", "s6x");
    replace_all(line, "seven", "s7n");
    replace_all(line, "eight", "e8t");
    replace_all(line, "nine", "n9e");
}

int main() {
    int ret = 0;

    std::string line;
    while (getline(std::cin, line)) {
        replace_by_numbers(line);

        int first = line.find_first_of(digits);
        if (first == std::string::npos) {
            continue;
        }

        int last = line.find_last_of(digits);
        if (last == std::string::npos) {
            continue;
        }

        std::string snum;
        snum += line[first];
        snum += line[last];

        ret += std::stoi(snum);
    }

    std::cout << ret << std::endl;

    return 0;
}
