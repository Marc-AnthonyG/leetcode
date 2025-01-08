const HOUR_ANGLE_UNIT = 360 / 12
const MINUTE_ANGLE_UNIT = 360 / 60

const HOUR_MINUTE_INFLUENCE_UNIT = HOUR_ANGLE_UNIT / 60

function angleClock(hour: number, minutes: number): number {
  const angleHourHand = hour * HOUR_ANGLE_UNIT + HOUR_MINUTE_INFLUENCE_UNIT * minutes
  const angleMinuteHand = minutes * MINUTE_ANGLE_UNIT

  const angle = Math.abs(angleHourHand - angleMinuteHand)

  return angle > 180 ? 360 - angle : angle
};


console.log(angleClock(12, 30)) // 165
