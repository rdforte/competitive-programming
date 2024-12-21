local function raindrops(n)
  local divisibleBy3 = n % 3 == 0
  local divisibleBy5 = n % 5 == 0
  local divisibleBy7 = n % 7 == 0

  local res = ""

  if divisibleBy3 then
    res = res .. "Pling"
  end

  if divisibleBy5 then
    res = res .. "Plang"
  end

  if divisibleBy7 then
    res = res .. "Plong"
  end

  if #res > 0 then
    return res
  end

  return tostring(n)
end

io.write(raindrops(1) .. "\n")
io.write(raindrops(3) .. "\n")
io.write(raindrops(5) .. "\n")
io.write(raindrops(7) .. "\n")
io.write(raindrops(21) .. "\n")
