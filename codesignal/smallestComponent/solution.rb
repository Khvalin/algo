def smallestComponent(m)
  n,w = m.size, m[0].size
  a, r = [[]], []
  c = -1
  v = lambda { |i, j, c|
    r[c] = 1 + (r[c] || 0)
    a[i][j] = c + 1
    for x in i-1..i+1
      for y in j-1..j+1
        m[x] && x >= 0 && y >= 0 && (x == i || y == j) &&
          (a[x] = a[x] || []) &&
          m[i][j] == m[x][y] &&
            !a[x][y] && v.(x, y, c) 
      end
    end
  }

  n.times{|i|
    w.times {|j|
      a[i][j] || v.(i, j, c+=1)
    }
  }
  
  #puts r.join ' '
  r.min
end

puts smallestComponent [
  "AABACAA", 
"ABBACCA", 
"ABBACCA", 
"ABBACCA", 
"ABBAACA", 
"AABBACA", 
"DAAAACA", 
"DACCCCA", 
"AAAAAAA"]


puts smallestComponent ["NNO", 
"ONO", 
"OON", 
"ONN"]