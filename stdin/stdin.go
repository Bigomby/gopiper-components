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
	"bufio"
	"os"

	"github.com/Bigomby/gopiper/component"
	"github.com/Bigomby/gopiper/messages"
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

// Factory is used to create instances of stdin component
type Factory struct{}

// NewFactory creates a new factory for stdin components
func NewFactory() component.Factory { return &Factory{} }

// Create is the method that creates instances of stdin component. It spawns
// a gorutine for reading data from stdin.
func (f *Factory) Create(out chan component.Message) component.Component {
	scanner := bufio.NewScanner(os.Stdin)

	go func() {
		for scanner.Scan() {
			msg := messages.NewBytesMessage()
			msg.SetData(scanner.Bytes())
			out <- msg
		}
	}()

	return &Component{}
}

// Destroy closes resources associated to the Factory
func (f Factory) Destroy() {}

// PoolSize returns the size of the worker pool
func (f Factory) PoolSize() int { return 1 }

// ChannelSize returns the size of the output channel
func (f Factory) ChannelSize() int { return 1 }

// SetAttribute allows to set attributes to the components created with the
// Create() method
func (f Factory) SetAttribute(string, interface{}) error { return nil }

///////////////
// Component //
///////////////

// Component is a component that reads data from stdin
type Component struct{}

// Handle does nothing since this component is the first on the pipeline
func (c *Component) Handle(component.Message) *component.Report { return nil }
