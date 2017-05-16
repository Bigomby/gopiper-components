// Copyright 2017 Diego Fernández Barrera
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"
	"strconv"

	"github.com/Bigomby/gopiper/component"
)

var (
	// Version of the component
	Version string
)

// Dummy main
func main() {}

/////////////
// Factory //
/////////////

// Factory is used to create instances of stdout component
type Factory struct {
	workers     int
	channelSize int
}

// NewFactory creates a new factory for stdout components
func NewFactory() component.Factory {
	return &Factory{channelSize: 1, workers: 1}
}

// Create is the method that creates instances of stdout component
func (f Factory) Create(out chan component.Message) component.Component {

	return &Component{}
}

// Destroy closes resources associated to the Factory
func (f Factory) Destroy() {}

// PoolSize returns the size of the worker pool
func (f Factory) PoolSize() int { return f.workers }

// ChannelSize returns the size of the output channel
func (f Factory) ChannelSize() int { return f.channelSize }

// SetAttribute allows to set attributes to the components created with the
// Create() method
func (f *Factory) SetAttribute(key string, value interface{}) error {
	switch key {

	case "workers":
		w, err := strconv.ParseInt(value.(string), 10, 32)
		if err != nil {
			return err
		}

		f.workers = int(w)

	case "channel_size":
		w, err := strconv.ParseInt(value.(string), 10, 32)
		if err != nil {
			return err
		}

		f.channelSize = int(w)
	}

	return nil
}

///////////////
// Component //
///////////////

// Component is a component that sends data to stdout
type Component struct{}

// Handle output the bytes received and write the data to stdout
func (c *Component) Handle(msg component.Message) *component.Report {
	os.Stdout.Write(msg.GetData().([]byte))
	os.Stdout.Write([]byte("\n"))
	return &component.Report{Status: component.Done}
}
