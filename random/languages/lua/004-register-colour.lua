local function color_code(color)
  local resistance = {
    black = 0,
    brown = 1,
    red = 2,
    orange = 3,
    yellow = 4,
    green = 5,
    blue = 6,
    violet = 7,
    grey = 8,
    white = 9,
  }

  return resistance[color]
end

io.write(tostring(color_code("black")) .. "\n")
io.write(tostring(color_code("brown")) .. "\n")
io.write(tostring(color_code("white")) .. "\n")
