class Solution {
    public int removeDuplicates(int[] nums) {
        int k = 1;
        
        for (int i = 1; i < nums.length; i++){
            //kita cek index bandingin index sebelumnya
            if (nums[i] != nums[i-1]){
                //kalo sama maka keluar loop, lnjt next index
                //kalo beda maka kita ubah array nums dengan posisi index k dengan value nums index i
                nums[k]= nums[i];
                k += 1;
            }
        }
        return k;
    }
}