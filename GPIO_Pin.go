// Copyright 2019 by mauro@ezplanet.org (Mauro Mozzarelli)
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
// this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
// this list of conditions and the following disclaimer in the documentation
// and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package pi_gpio

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

type GPIO_Pin struct {
	Name string
	Path string
}

func (r GPIO_Pin) BasePath() string {
	return GPIO_GPIO + r.Name
}

func (r GPIO_Pin) write(where, what string) GPIO_Pin {
	filename := r.Path + "/" + where
	err := ioutil.WriteFile(filename, []byte(what), 0666)
	if err != nil {
		log.Println(err)
	}
	return r
}

func (r GPIO_Pin) read(where string) string {
	filename := r.Path + "/" + where
	status, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
	}
	return strings.TrimSuffix(string(status), "\n")
}

func (r GPIO_Pin) Export() GPIO_Pin {
	r.Path = r.BasePath()
	if _, err := os.Stat(r.Path); os.IsNotExist(err) {
		log.Printf("pin directory not present, error: %s", err)
		// export gpio pin
		err := ioutil.WriteFile(GPIO_EXPORT, []byte(r.Name), 0666)
		if err != nil {
			log.Println(err)
		} else {
			time.Sleep(2000 * time.Millisecond)
			log.Printf("GPIO.Pin %s exported, file: %s", r.Name, GPIO_EXPORT)
		}
	} else {
		log.Printf("path %s already exists", r.Path)
	}
	return GPIO_Pin{}
}

func (r GPIO_Pin) Unexport() GPIO_Pin {
	r.Path = r.BasePath()
	if _, err := os.Stat(r.Path); os.IsNotExist(err) {
		log.Printf("pin directory not present, error: %s", err)
	} else {
		err := ioutil.WriteFile(GPIO_UNEXPORT, []byte(r.Name), 0666)
		if err != nil {
			log.Println(err)
		} else {
			time.Sleep(2000 * time.Millisecond)
			log.Printf("GPIO.Pin %s unexported, file: %s", r.Name, GPIO_UNEXPORT)
		}
	}
	return GPIO_Pin{}
}
func (r GPIO_Pin) Input() GPIO_Pin {
	log.Printf("PIN %s INPUT", r.Name)
	return r.write(DIRECTION, IN)
}

func (r GPIO_Pin) Output() GPIO_Pin {
	log.Printf("PIN %s OUTPUT", r.Name)
	return r.write(DIRECTION, OUT)
}

func (r GPIO_Pin) High() GPIO_Pin {
	log.Printf("PIN %s HIGH", r.Name)
	return r.write(VALUE, ON)
}

func (r GPIO_Pin) Low() GPIO_Pin {
	log.Printf("PIN %s LOW", r.Name)
	return r.write(VALUE, OFF)
}

func (r GPIO_Pin) Status() string {
	return r.read(VALUE)
}
