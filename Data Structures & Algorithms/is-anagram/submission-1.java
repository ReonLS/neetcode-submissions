class Solution {
    public boolean isAnagram(String s, String t) {
        int[] count = new int[26];

        //validasi 1, anagram pasti punya length sama stringnya
        if (s.length() != t.length()){
            return false;
        }

        //validasi 2, check mereka contain same chars gk
        for (int i = 0; i < s.length(); i++){
            count[s.charAt(i) - 'a']++;
            count[t.charAt(i) - 'a']--;
        }

        //sekarang check, harusnya kalo anagram maka value count semuanya '0'
        for (int k = 0; k < 26; k++){
            if (count[k] != 0){
                return false;
            }
        }
        return true;
    }
}
