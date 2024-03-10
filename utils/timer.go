package utils

import rl "github.com/gen2brain/raylib-go/raylib"

type Timer struct {
	life_time float32
}

func NewTimer() Timer {
	return Timer{
		life_time: 0.0,
	}
}

func (t *Timer) Start(time float32) {
	t.life_time = time
}

func (t *Timer) Update() {
	if t.life_time > 0 {
		t.life_time -= rl.GetFrameTime()
	}
}
func (t *Timer) IsFinished() bool {
	return t.life_time <= 0
}

/*


#ifndef TIMER_H
#define TIMER_H

#include "raylib.h"

class Timer {
    public:
        float life_time;

    static Timer New();

    void start(float life_time);
    void update();
    bool isFinished();




};

#endif

#include "Timer.h"
#include "raylib.h"
#include "raymath.h"

Timer Timer::New(){
    Timer timer;
    return timer;
}

void Timer::start(float life_time) {
    this->life_time = life_time;
}

void Timer::update(){
    if (this->life_time > 0){
        this->life_time -= GetFrameTime();
    }
}


bool Timer::isFinished(){
    return (life_time <= 0) ;
}
*/
