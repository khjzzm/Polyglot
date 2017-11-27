puts("아이디를 입력해주세요")
in_id = gets.chomp()

puts("비밀번호를 입력해주세요")
in_password = gets.chomp()

real_id = "kim"
real_passowrd = "1234"

if(in_id == real_id)
  if(in_password == real_passowrd)
    puts(in_id+"님 환영합니다")
  else
    puts(real_passowrd+"는 잘못된 비밀번호입니다")
  end
else
  puts(in_id+"는 없는 아이디입니다.")
end

=begin
권장하는 주석처리 방법이 아니다
=end

# 권장하는 주석