start_time="$(date -u +%s.%N)"
echo "Starting tests..."
echo

echo go run .  
echo
go run . 
echo

echo go run . --align=left "€ä" standard; 
echo
go run . --align=left "€ä" standard;  
echo

echo go run . --align=foobar "foo bar zoo"
echo
go run . --align=foobar "foo bar zoo"
echo

echo go run . --align"\033[41m \033[0m"right something standard
echo
go run . --align right something standard

echo go run . "Hello" shadow
echo
go run . "Hello" shadow

echo go run . --align=right left standard
echo
go run . --align=right left standard

echo go run . --align=left right standard
echo
go run . --align=left right standard

echo go run . --align=center hello shadow
echo
go run . --align=center hello shadow

echo go run . --align=justify "1 Two 4" shadow
echo
go run . --align=justify "1 Two 4" shadow

echo go run . --align=right 23/32 standard
echo
go run . --align=right 23/32 standard

echo go run . --align=right ABCabc123 tinkertoy
echo
go run . --align=right ABCabc123 tinkertoy

echo go run . --align=center "#$%&\"" tinkertoy
echo
go run . --align=center "#$%&\"" tinkertoy

echo go run . --align=left "23Hello World!" standard
echo
go run . --align=left "23Hello World!" standard

echo go run . --align=justify "HELLO there HOW are YOU?!" tinkertoy
echo
go run . --align=justify "HELLO there HOW are YOU?!" tinkertoy

echo go run . --align=right "a -> A b -> B c -> C" shadow
echo
go run . --align=right "a -> A b -> B c -> C" shadow

echo go run . --align=right abcd shadow
echo
go run . --align=right abcd shadow

echo go run . --align=center ola standard
echo
go run . --align=center ola standard

echo go run . --align=right "oKcZVuF" standard
echo
go run . --align=right "oKcZVuF" standard

echo go run . --align=center "0 9 4 qkv" standard
echo
go run . --align=center "0 9 4 qkv" standard 

echo go run . --align=justify "!@#$ %^ &*()" standard  
echo
go run . --align=justify "!@#$ %^ &*()" standard 

echo go run . --align=left "R9c E5x p8f W2n" standard  
echo
go run . --align=left "R9c E5x p8f W2n" standard

end_time="$(date -u +%s.%N)"
echo "Tests finished!"
elapsed="$(awk "BEGIN { print $end_time - $start_time }")"
echo "Total of $elapsed seconds elapsed for processing"