// Code generated by gop (Go+); DO NOT EDIT.

package main

import "github.com/goplus/spx"

const _ = true
const winscore = 500
const (
	msgBattle   = "battle"
	msgKillAll  = "kill all"
	msgYouWin   = "you win"
	msgGameOver = "game over"
	msgRestart  = "restart"
)

type Backdrop struct {
	spx.Sprite
	*Game
}
type Bomb struct {
	spx.Sprite
	*Game
}
type Bullet struct {
	spx.Sprite
	*Game
}
type GameLogo struct {
	spx.Sprite
	*Game
}
type GameOver struct {
	spx.Sprite
	*Game
}
type GameStart struct {
	spx.Sprite
	*Game
}
type HugeEnemy struct {
	spx.Sprite
	*Game
	life int
}
type MiddleEnemy struct {
	spx.Sprite
	*Game
	life int
}
type MyAircraft struct {
	spx.Sprite
	*Game
	life int
}
type Restart struct {
	spx.Sprite
	*Game
}
type SmallEnemy struct {
	spx.Sprite
	*Game
	life int
}
type TextIntro struct {
	spx.Sprite
	*Game
}
type YouWin struct {
	spx.Sprite
	*Game
}
type Game struct {
	spx.Game
	Backdrop    Backdrop
	Bomb        Bomb
	TextIntro   TextIntro
	GameLogo    GameLogo
	GameStart   GameStart
	MyAircraft  MyAircraft
	MiddleEnemy MiddleEnemy
	HugeEnemy   HugeEnemy
	Bullet      Bullet
	SmallEnemy  SmallEnemy
	GameOver    GameOver
	YouWin      YouWin
	Restart     Restart
	bombs       int
	score       int
}
//line main.spx:32:1
func (this *Game) addScore(val int) {
//line main.spx:33:1
	this.score += val
//line main.spx:34:1
	if this.score > winscore {
//line main.spx:35:1
		this.Broadcast__0(msgYouWin)
	}
}
//line main.spx:39:1
func (this *Game) restart() {
//line main.spx:40:1
	spx.Gopt_Game_Reload(this, "index.json")
//line main.spx:41:1
	this.GameStart.Hide()
//line main.spx:42:1
	this.Broadcast__0(msgBattle)
}
func (this *Game) MainEntry() {
}
func (this *Game) Main() {
	spx.Gopt_Game_Main(this, new(Backdrop), new(Bomb), new(Bullet), new(GameLogo), new(GameOver), new(GameStart), new(HugeEnemy), new(MiddleEnemy), new(MyAircraft), new(Restart), new(SmallEnemy), new(TextIntro), new(YouWin))
}
func (this *Backdrop) Classfname() string {
	return "Backdrop"
}
func (this *Backdrop) Main() {
}
//line Bomb.spx:1
func (this *Bomb) Main() {
//line Bomb.spx:1:1
	this.OnTouched__1(func() {
//line Bomb.spx:2:1
		this.Destroy()
	})
//line Bomb.spx:5:1
	this.OnCloned__1(func() {
//line Bomb.spx:6:1
		this.SetXYpos(spx.Rand__0(-90, 90), 237)
//line Bomb.spx:7:1
		this.Show()
//line Bomb.spx:8:1
		for {
			spx.Sched()
//line Bomb.spx:9:1
			this.ChangeYpos(-5)
//line Bomb.spx:10:1
			if this.Ypos() < -170 {
//line Bomb.spx:11:1
				this.Destroy()
			}
//line Bomb.spx:13:1
			this.Wait(0.05)
		}
	})
//line Bomb.spx:17:1
	this.OnMsg__1(msgBattle, func() {
//line Bomb.spx:18:1
		this.bombs = 0
//line Bomb.spx:19:1
		for
//line Bomb.spx:19:1
		i := 0; i < 50;
//line Bomb.spx:19:1
		i++ {
			spx.Sched()
//line Bomb.spx:20:1
			this.Wait(20)
//line Bomb.spx:21:1
			spx.Gopt_Sprite_Clone__0(this)
		}
	})
//line Bomb.spx:25:1
	this.OnMsg__1(msgGameOver, func() {
//line Bomb.spx:26:1
		this.Stop(spx.OtherScriptsInSprite)
	})
//line Bomb.spx:28:1
	this.OnMsg__1(msgYouWin, func() {
//line Bomb.spx:29:1
		this.Stop(spx.OtherScriptsInSprite)
	})
}
func (this *Bomb) Classfname() string {
	return "Bomb"
}
//line Bullet.spx:1
func (this *Bullet) Main() {
//line Bullet.spx:1:1
	this.OnTouched__1(func() {
//line Bullet.spx:2:1
		this.Destroy()
	})
//line Bullet.spx:5:1
	this.OnCloned__1(func() {
//line Bullet.spx:6:1
		this.SetXYpos(this.MyAircraft.Xpos(), this.MyAircraft.Ypos()+5)
//line Bullet.spx:7:1
		this.Show()
//line Bullet.spx:8:1
		for {
			spx.Sched()
//line Bullet.spx:9:1
			this.Step__0(10)
//line Bullet.spx:10:1
			this.Wait(0.04)
//line Bullet.spx:11:1
			if this.Touching(spx.Edge) {
//line Bullet.spx:12:1
				this.Destroy()
			}
		}
	})
//line Bullet.spx:17:1
	this.OnMsg__1(msgGameOver, func() {
//line Bullet.spx:18:1
		this.Stop(spx.OtherScriptsInSprite)
	})
//line Bullet.spx:20:1
	this.OnMsg__1(msgYouWin, func() {
//line Bullet.spx:21:1
		this.Stop(spx.OtherScriptsInSprite)
	})
}
func (this *Bullet) Classfname() string {
	return "Bullet"
}
//line GameLogo.spx:1
func (this *GameLogo) Main() {
//line GameLogo.spx:1:1
	this.OnMsg__1(msgBattle, func() {
//line GameLogo.spx:2:1
		this.Hide()
	})
}
func (this *GameLogo) Classfname() string {
	return "GameLogo"
}
//line GameOver.spx:1
func (this *GameOver) Main() {
//line GameOver.spx:1:1
	this.OnMsg__1(msgBattle, func() {
//line GameOver.spx:2:1
		this.Hide()
	})
//line GameOver.spx:5:1
	this.OnMsg__1(msgGameOver, func() {
//line GameOver.spx:6:1
		this.Show()
//line GameOver.spx:7:1
		this.Wait(0.8)
//line GameOver.spx:8:1
		this.Broadcast__0(msgRestart)
	})
}
func (this *GameOver) Classfname() string {
	return "GameOver"
}
//line GameStart.spx:1
func (this *GameStart) Main() {
//line GameStart.spx:1:1
	this.OnClick(func() {
//line GameStart.spx:2:1
		this.Hide()
//line GameStart.spx:3:1
		this.Broadcast__0(msgBattle)
	})
}
func (this *GameStart) Classfname() string {
	return "GameStart"
}
//line HugeEnemy.spx:5
func (this *HugeEnemy) Main() {
//line HugeEnemy.spx:5:1
	this.OnMsg__1(msgBattle, func() {
//line HugeEnemy.spx:6:1
		for {
			spx.Sched()
//line HugeEnemy.spx:7:1
			this.Wait(25)
//line HugeEnemy.spx:8:1
			spx.Gopt_Sprite_Clone__0(this)
		}
	})
//line HugeEnemy.spx:12:1
	this.OnCloned__1(func() {
//line HugeEnemy.spx:13:1
		this.life = 50
//line HugeEnemy.spx:14:1
		this.SetCostume(0)
//line HugeEnemy.spx:15:1
		this.SetXYpos(spx.Rand__0(-110, 110), 237)
//line HugeEnemy.spx:16:1
		this.Show()
//line HugeEnemy.spx:17:1
		for {
			spx.Sched()
//line HugeEnemy.spx:18:1
			this.ChangeYpos(-2)
//line HugeEnemy.spx:19:1
			if this.Ypos() < -200 {
//line HugeEnemy.spx:20:1
				this.Destroy()
			}
//line HugeEnemy.spx:22:1
			if this.Touching("Bullet") {
//line HugeEnemy.spx:23:1
				this.life--
//line HugeEnemy.spx:24:1
				if this.life == 0 {
//line HugeEnemy.spx:25:1
					this.addScore(50)
//line HugeEnemy.spx:26:1
					this.Die()
				}
			}
//line HugeEnemy.spx:29:1
			this.Wait(0.05)
		}
	})
//line HugeEnemy.spx:33:1
	this.OnMsg__1(msgKillAll, func() {
//line HugeEnemy.spx:34:1
		this.addScore(50)
//line HugeEnemy.spx:35:1
		this.Die()
	})
//line HugeEnemy.spx:38:1
	this.OnMsg__1(msgGameOver, func() {
//line HugeEnemy.spx:39:1
		this.Stop(spx.OtherScriptsInSprite)
	})
//line HugeEnemy.spx:41:1
	this.OnMsg__1(msgYouWin, func() {
//line HugeEnemy.spx:42:1
		this.Stop(spx.OtherScriptsInSprite)
	})
}
func (this *HugeEnemy) Classfname() string {
	return "HugeEnemy"
}
//line MiddleEnemy.spx:5
func (this *MiddleEnemy) Main() {
//line MiddleEnemy.spx:5:1
	this.OnMsg__1(msgBattle, func() {
//line MiddleEnemy.spx:6:1
		for {
			spx.Sched()
//line MiddleEnemy.spx:7:1
			this.Wait(2)
//line MiddleEnemy.spx:8:1
			spx.Gopt_Sprite_Clone__0(this)
		}
	})
//line MiddleEnemy.spx:12:1
	this.OnCloned__1(func() {
//line MiddleEnemy.spx:13:1
		this.life = 8
//line MiddleEnemy.spx:14:1
		this.SetCostume(0)
//line MiddleEnemy.spx:15:1
		this.SetXYpos(spx.Rand__0(-131, 131), 237)
//line MiddleEnemy.spx:16:1
		this.Show()
//line MiddleEnemy.spx:17:1
		for {
			spx.Sched()
//line MiddleEnemy.spx:18:1
			this.ChangeYpos(-5)
//line MiddleEnemy.spx:19:1
			if this.Ypos() < -170 {
//line MiddleEnemy.spx:20:1
				this.Destroy()
			}
//line MiddleEnemy.spx:22:1
			if this.Touching("Bullet") {
//line MiddleEnemy.spx:23:1
				this.life--
//line MiddleEnemy.spx:24:1
				if this.life == 0 {
//line MiddleEnemy.spx:25:1
					this.addScore(10)
//line MiddleEnemy.spx:26:1
					this.Die()
				}
			}
//line MiddleEnemy.spx:29:1
			this.Wait(0.05)
		}
	})
//line MiddleEnemy.spx:33:1
	this.OnMsg__1(msgKillAll, func() {
//line MiddleEnemy.spx:34:1
		this.addScore(10)
//line MiddleEnemy.spx:35:1
		this.Die()
	})
//line MiddleEnemy.spx:38:1
	this.OnMsg__1(msgGameOver, func() {
//line MiddleEnemy.spx:39:1
		this.Stop(spx.OtherScriptsInSprite)
	})
//line MiddleEnemy.spx:42:1
	this.OnMsg__1(msgYouWin, func() {
//line MiddleEnemy.spx:43:1
		this.Stop(spx.OtherScriptsInSprite)
	})
}
func (this *MiddleEnemy) Classfname() string {
	return "MiddleEnemy"
}
//line MyAircraft.spx:5
func (this *MyAircraft) Main() {
//line MyAircraft.spx:5:1
	this.OnMsg__1(msgBattle, func() {
//line MyAircraft.spx:6:1
		this.bombs = 0
//line MyAircraft.spx:7:1
		this.score = 0
//line MyAircraft.spx:8:1
		this.life = 1
//line MyAircraft.spx:9:1
		for this.life > 0 {
			spx.Sched()
//line MyAircraft.spx:10:1
			this.Wait(0.1)
//line MyAircraft.spx:11:1
			spx.Gopt_Sprite_Clone__0(&this.Bullet)
		}
	})
//line MyAircraft.spx:15:1
	this.OnMsg__1(msgBattle, func() {
//line MyAircraft.spx:16:1
		this.Show()
//line MyAircraft.spx:17:1
		this.SetXYpos(0, 0)
//line MyAircraft.spx:18:1
		this.score = 0
//line MyAircraft.spx:19:1
		for {
			spx.Sched()
//line MyAircraft.spx:20:1
			this.Wait(0.01)
//line MyAircraft.spx:21:1
			dis := this.DistanceTo(spx.Mouse)
//line MyAircraft.spx:22:1
			xx := this.Xpos()
//line MyAircraft.spx:23:1
			yy := this.Ypos()
//line MyAircraft.spx:24:1
			speed := 3.0
//line MyAircraft.spx:25:1
			planeX := 0.0
//line MyAircraft.spx:26:1
			planeY := 0.0
//line MyAircraft.spx:27:1
			if dis > speed {
//line MyAircraft.spx:28:1
				planeX = xx + (this.MouseX()-xx)/dis*speed
//line MyAircraft.spx:29:1
				planeY = yy + (this.MouseY()-yy)/dis*speed
			} else {
//line MyAircraft.spx:31:1
				planeX = this.MouseX()
//line MyAircraft.spx:32:1
				planeY = this.MouseY()
			}
//line MyAircraft.spx:34:1
			if planeX < -131 {
//line MyAircraft.spx:35:1
				planeX = -131
			} else
//line MyAircraft.spx:36:1
			if planeX > 131 {
//line MyAircraft.spx:37:1
				planeX = 131
			}
//line MyAircraft.spx:39:1
			if planeY < -170 {
//line MyAircraft.spx:40:1
				planeY = -170
			}
//line MyAircraft.spx:42:1
			this.SetXYpos(planeX, planeY)
//line MyAircraft.spx:43:1
			if this.Touching("HugeEnemy") || this.Touching("SmallEnemy") || this.Touching("MiddleEnemy") {
//line MyAircraft.spx:44:1
				this.life--
//line MyAircraft.spx:45:1
				this.Broadcast__0(msgGameOver)
//line MyAircraft.spx:46:1
				this.Die()
			}
//line MyAircraft.spx:48:1
			if this.Touching("Bomb") {
//line MyAircraft.spx:49:1
				this.bombs++
			}
		}
	})
//line MyAircraft.spx:54:1
	this.OnKey__0(spx.KeySpace, func() {
//line MyAircraft.spx:55:1
		if this.bombs > 0 {
//line MyAircraft.spx:56:1
			this.bombs--
//line MyAircraft.spx:57:1
			this.Broadcast__0(msgKillAll)
		}
	})
//line MyAircraft.spx:60:1
	this.OnMsg__1(msgYouWin, func() {
//line MyAircraft.spx:61:1
		this.Stop(spx.OtherScriptsInSprite)
	})
}
func (this *MyAircraft) Classfname() string {
	return "MyAircraft"
}
//line Restart.spx:1
func (this *Restart) Main() {
//line Restart.spx:1:1
	this.OnMsg__1(msgRestart, func() {
//line Restart.spx:2:1
		this.Wait(0.3)
//line Restart.spx:3:1
		this.GotoFront()
//line Restart.spx:4:1
		this.Show()
//line Restart.spx:5:1
		for {
			spx.Sched()
//line Restart.spx:6:1
			if this.Touching(spx.Mouse) {
//line Restart.spx:7:1
				this.SetYpos(-90)
			} else {
//line Restart.spx:9:1
				this.SetYpos(-95)
			}
//line Restart.spx:11:1
			this.Wait(0.1)
		}
	})
//line Restart.spx:15:1
	this.OnClick(func() {
//line Restart.spx:16:1
		this.Hide()
//line Restart.spx:17:1
		this.Wait(0.3)
//line Restart.spx:18:1
		this.restart()
	})
}
func (this *Restart) Classfname() string {
	return "Restart"
}
//line SmallEnemy.spx:5
func (this *SmallEnemy) Main() {
//line SmallEnemy.spx:5:1
	this.OnMsg__1(msgBattle, func() {
//line SmallEnemy.spx:6:1
		for {
			spx.Sched()
//line SmallEnemy.spx:7:1
			this.Wait(1)
//line SmallEnemy.spx:8:1
			spx.Gopt_Sprite_Clone__0(this)
		}
	})
//line SmallEnemy.spx:12:1
	this.OnCloned__1(func() {
//line SmallEnemy.spx:13:1
		this.life = 1
//line SmallEnemy.spx:14:1
		this.SetCostume(0)
//line SmallEnemy.spx:15:1
		this.SetXYpos(spx.Rand__0(-131, 131), 237)
//line SmallEnemy.spx:16:1
		this.Show()
//line SmallEnemy.spx:17:1
		for {
			spx.Sched()
//line SmallEnemy.spx:18:1
			this.ChangeYpos(-8)
//line SmallEnemy.spx:19:1
			if this.Ypos() < -170 {
//line SmallEnemy.spx:20:1
				this.Destroy()
			}
//line SmallEnemy.spx:22:1
			if this.Touching("Bullet") {
//line SmallEnemy.spx:23:1
				this.life--
//line SmallEnemy.spx:24:1
				if this.life == 0 {
//line SmallEnemy.spx:25:1
					this.addScore(5)
//line SmallEnemy.spx:26:1
					this.Die()
				}
			}
//line SmallEnemy.spx:29:1
			this.Wait(0.05)
		}
	})
//line SmallEnemy.spx:33:1
	this.OnMsg__1(msgKillAll, func() {
//line SmallEnemy.spx:34:1
		this.addScore(5)
//line SmallEnemy.spx:35:1
		this.Die()
	})
//line SmallEnemy.spx:38:1
	this.OnMsg__1(msgGameOver, func() {
//line SmallEnemy.spx:39:1
		this.Stop(spx.OtherScriptsInSprite)
	})
//line SmallEnemy.spx:41:1
	this.OnMsg__1(msgYouWin, func() {
//line SmallEnemy.spx:42:1
		this.Stop(spx.OtherScriptsInSprite)
	})
}
func (this *SmallEnemy) Classfname() string {
	return "SmallEnemy"
}
//line TextIntro.spx:1
func (this *TextIntro) Main() {
//line TextIntro.spx:1:1
	this.OnMsg__1(msgBattle, func() {
//line TextIntro.spx:2:1
		this.Hide()
	})
}
func (this *TextIntro) Classfname() string {
	return "TextIntro"
}
//line YouWin.spx:1
func (this *YouWin) Main() {
//line YouWin.spx:1:1
	this.OnMsg__1(msgYouWin, func() {
//line YouWin.spx:2:1
		this.Show()
//line YouWin.spx:3:1
		this.Wait(0.8)
//line YouWin.spx:4:1
		this.Broadcast__0(msgRestart)
	})
}
func (this *YouWin) Classfname() string {
	return "YouWin"
}
func main() {
	new(Game).Main()
}
