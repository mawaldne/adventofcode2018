file = File.open("input.txt", "rb")
licence_file = file.read.split.map(&:to_i)

header_stack = []
meta_list = []

children, meta = licence_file.slice!(0, 2)
header_stack << [children, meta]

until licence_file.empty?
  # process stack
  children, meta = header_stack.shift
  if children == 0 then
    meta.times { meta_list << licence_file.shift }
  else
    children -= 1
    header_stack.unshift([children, meta])
    header_stack.unshift(licence_file.slice!(0, 2))
  end
end

pp meta_list
pp meta_list.sum

