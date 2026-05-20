class Solution {
    int[] indices = new int[2];

    public int[] twoSum(int[] nums, int target) {
        for (int x = 0; x < nums.length; x++)
        {
            for (int y = x+1; y < nums.length; y++)
            {
                if (nums[y] == target-nums[x])
                {
                    indices[0] = x;
                    indices[1] = y;
                    break;
                }
            }
        }
        return indices;
    }
}
