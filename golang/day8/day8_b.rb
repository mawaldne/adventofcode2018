file = File.open("input.txt", "rb")
licence_file = file.read.split.map(&:to_i)

header_stack = []

children, meta = licence_file.slice!(0, 2)
# children total, meta, children value array
header_stack << [children, meta, []]

until licence_file.empty?
  # process stack
  children, meta, child_values = header_stack.shift
  if children == 0 then
    meta_list = []
    meta.times { meta_list << licence_file.shift }

    child_total = 0
    if child_values.empty?
      child_total = meta_list.sum
    else
      child_total = meta_list.map { |i| child_values[i - 1] }.compact.sum
    end

    if header_stack.any? then
      children, meta, child_values = header_stack.shift
      header_stack.unshift([children, meta, child_values << child_total])
    else
      pp "Answer = #{child_total}"
    end
  else
    children -= 1
    header_stack.unshift([children, meta, child_values])
    children, meta = licence_file.slice!(0, 2)
    header_stack.unshift([children, meta, []])
  end
end
