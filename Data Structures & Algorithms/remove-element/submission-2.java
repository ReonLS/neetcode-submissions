class Solution {
    int k = 0;

    public int removeElement(int[] nums, int val)
    {
        //pointer1
        for (int i = 0; i < nums.length; i++)
        {
            //kalo val, maka candidate untuk change place
            if (nums[i] == val)
            {
                for (int j = i+1; j < nums.length; j++)
                {
                    //ubah posisi apabila ketemu yg value bukan val
                    if(nums[j] != val)
                    {
                        nums[i] = nums[j];
                        nums[j] = val;
                        k++;
                        break;
                    }
                }
            }
            else {
                k++;
            }
        }
        return k;
    }
}