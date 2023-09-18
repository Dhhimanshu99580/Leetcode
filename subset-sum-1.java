class Solution{
    ArrayList<Integer> subsetSums(ArrayList<Integer> arr, int N){
        ArrayList<Integer>sumSet=new ArrayList<Integer>();
        findSum(arr,sumSet,N,0,0);
        return sumSet;
    }
    public void findSum(ArrayList<Integer>arr,ArrayList<Integer>sumSet,int n,int i, int sum) {
        if(i==n) {
            sumSet.add(sum);
            return;
        }
        sum+=arr.get(i);
        i++;
        findSum(arr,sumSet,n,i,sum);
        sum-=arr.get(i-1);
       
        findSum(arr,sumSet,n,i,sum);
    }
}