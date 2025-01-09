function subsets(nums: number[]): number[][] {
  const ans: number[][] = [];

  const createSubset = (currentArray: number[], index: number) => {
    if (index === nums.length - 1) {
      ans.push(currentArray)
      ans.push([...currentArray, nums[index]])
    } else {
      createSubset(currentArray, index + 1)
      createSubset([...currentArray, nums[index]], index + 1)
    }
  }

  createSubset([], 0)

  return ans
}


console.log(subsets([1, 2, 3]))
