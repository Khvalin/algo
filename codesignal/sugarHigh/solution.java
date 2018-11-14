 static int[] sugarHigh(int[] candies, int threshold) {
        int[] sorted = candies.clone();
        Arrays.sort(sorted);
        System.out.println(Arrays.toString(sorted));

        HashMap<Integer, Integer> map = new HashMap<>();
        int n = 0;
        while (n < sorted.length && sorted[n] <= threshold) {
            int s = sorted[n];
            threshold -= s;
            int c = map.containsKey(s)? map.get(s):0;
            map.put(s, c + 1);
            n++;
        }

        ArrayList<Integer> res = new ArrayList<>();
        for (int i = 0; i < candies.length;i++) {
            int candy = candies[i];
            int c = map.containsKey(candy)? map.get(candy):0;
            if (c > 0) {
                res.add(i);
                map.put(candy, c - 1);
            }
        }

        int[] array = res.stream().mapToInt(Integer::intValue).toArray();

        return  array;
    }