def multipleChoiceCheaters(a, b, k)
  s, n = 7, k.size
  res1 = 0
  acc, l = 0, 0
  n.times{|i| 
    prev = (i >= s && (a[i-s] != k[i-s] && b[i-s] == a[i-s])) ? 1 : 0
    cur = (a[i] != k[i] && b[i] == a[i] ? 1 : 0)
    acc += cur - prev

    if (i >=l + s - 1 && acc >= 4)
      res1 += s
      l = i+1
    end
  }

  #puts "#{l} #{res1} #{n-res1}"

  res1 > n-res1
end

puts multipleChoiceCheaters "DDDDBBBCCBBBAAAA", "DDDDBBBCCBBBAAAA", "AAAABBBCCBBBDDDD"

#puts multipleChoiceCheaters "ABCDAABCBACAB", "ABCDAABCBACAB", "ABCDAABCCDBDA"
