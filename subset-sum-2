class Solution {
    public List<List<Integer>> subsetsWithDup(int[] nums) {
        List<List<Integer>> lol = new ArrayList<List<Integer>>();
        List<Integer> list = new ArrayList<>();
        Arrays.sort(nums);
        fun(0, lol, list, nums);
        return lol;
    }

    public void fun(int index, List<List<Integer>> lol, List<Integer> list, int arr[]) {
        if (index >= arr.length) {
            if (!lol.contains(new ArrayList<>(list))) {
                lol.add(new ArrayList<>(list));
                return;
            } else {
                return;
            }
        }
        list.add(arr[index]);
        fun(index + 1, lol, list, arr);
        list.remove(list.size() - 1);
        fun(index + 1, lol, list, arr);
        return;
    }
}