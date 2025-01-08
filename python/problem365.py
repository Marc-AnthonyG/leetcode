def canMeasureWater(jug1Capacity: int, jug2Capacity: int, targetCapacity: int) -> bool:
        if jug1Capacity + jug2Capacity != targetCapacity and targetCapacity > jug2Capacity or  targetCapacity > jug1Capacity:
            print("here")
            return False
        elif jug1Capacity == targetCapacity or jug2Capacity == targetCapacity:
            return True 
        

        smallJug = jug1Capacity
        bigJug = jug2Capacity
        if jug1Capacity > jug2Capacity:
            smallJug = jug2Capacity
            bigJug = jug1Capacity

        print(smallJug, bigJug)

        if targetCapacity == bigJug - smallJug:
            return True
        elif targetCapacity % smallJug == smallJug - bigJug % smallJug:
            return True

        return False

canMeasureWater(3,5,4)
