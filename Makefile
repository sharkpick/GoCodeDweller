CXX:=c++ -std=c++2a -Wall -Wextra -pedantic

mangler/mangler.hpp:
	cp -r CodeDweller/mangler.*pp mangler/

clean: cleanup_mangler

cleanup_mangler:
	rm -f mangler/mangler.hpp mangler/mangler.cpp