import React, { useState, useEffect } from "react";

interface TimerProps {
  initialMinutes: number;
  initialSeconds: number;
  onTimerEnd: () => void;
}

const Timer: React.FC<TimerProps> = ({
  initialMinutes,
  initialSeconds,
  onTimerEnd,
}) => {
  const [minutes, setMinutes] = useState(initialMinutes);
  const [seconds, setSeconds] = useState(initialSeconds);

  useEffect(() => {
    const timer = setInterval(() => {
      if (seconds > 0) {
        setSeconds(seconds - 1);
      }
      if (seconds === 0) {
        if (minutes === 0) {
          clearInterval(timer);
          onTimerEnd();
        } else {
          setMinutes(minutes - 1);
          setSeconds(59);
        }
      }
    }, 1000);

    return () => clearInterval(timer);
  }, [minutes, seconds, onTimerEnd]);

  return (
    <div className="flex items-center justify-center">
      <div className="bg-dnRed text-white rounded-lg p-6 shadow-lg">
        <div className="text-4xl font-bold font-mono">
          {String(minutes).padStart(2, "0")}:{String(seconds).padStart(2, "0")}
        </div>
      </div>
    </div>
  );
};

export default Timer;
