local leap_year = function(number)
  local isDivisibleBy4 = number % 4 == 0
  local isDivisibleBy100 = number % 100 == 0
  local isDivisibleBy400 = number % 400 == 0
  if isDivisibleBy100 then
    return isDivisibleBy400
  end
  return isDivisibleBy4
end

io.write(tostring(leap_year(1996) == true) .. "\n")
io.write(tostring(leap_year(1997) == false) .. "\n")
io.write(tostring(leap_year(1900) == false) .. "\n")
io.write(tostring(leap_year(2000) == true) .. "\n")
