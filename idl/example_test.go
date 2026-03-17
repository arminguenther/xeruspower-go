// SPDX-License-Identifier: BSD-3-Clause
// Copyright 2026 Raritan Inc. All rights reserved.

package idl

import "fmt"

func ExampleParseTypeCode() {
	code, ok := ParseTypeCode("foo.Foo:1.2.3")
	fmt.Println(code.Name, code.Major, code.Submajor, code.Minor, ok)
	// Output: foo.Foo 1 2 3 true
}

func ExampleTypeCode_String() {
	code, _ := ParseTypeCode("foo.Foo:1.2.3")
	fmt.Println(code.String())
	// Output: foo.Foo:1.2.3
}

func ExampleTypeCode_StringNoVersions() {
	code, _ := ParseTypeCode("foo.Foo_1_2_3.Bar:4.5.6")
	fmt.Println(code.StringNoVersions())
	// Output: foo.Foo.Bar
}

func ExampleTypeCode_IsCompatible() {
	server, _ := ParseTypeCode("foo.Abc_1_0_1.Foo_2_2_2.Bar:5.5.5")
	clients := []string{
		"foo.Abc_1_0_1.Foo_2_2_2.Bar:5.5.5", // equal
		"foo.Abc_1_0_1.Foo_2_2_2.Boo:5.5.5", // Boo != Bar
		"foo.Abc_1_0_1.Foo_2_2_2.Bar:6.5.5", // higher Major
		"foo.Abc_1_0_1.Foo_2_2_2.Bar:4.5.5", // lower Major
		"foo.Abc_1_0_1.Foo_2_2_2.Bar:5.6.5", // call compatible
		"foo.Abc_1_0_1.Foo_2_2_2.Bar:5.4.5", // call compatible
		"foo.Abc_1_0_1.Foo_2_2_2.Bar:5.5.6", // higher Minor
		"foo.Abc_1_0_1.Foo_2_2_2.Bar:5.5.4", // backward compatible
		"foo.Abc_1_0_1.Fuu_2_2_2.Bar:5.5.5", // Fuu.Bar != Foo.Bar
		"foo.Abc_1_0_1.Foo_3_2_2.Bar:5.5.5", // Foo incompatible
		"foo.Abc_1_0_1.Foo_2_1_2.Bar:5.5.5", // call compatible
		"foo.Abc_1_0_1.Foo_2_2_3.Bar:5.5.5", // Foo incompatible
		"foo.Abc_1_0_1.Foo_2_2_1.Bar:5.5.5", // backward compatible
		"foo.Abc.Foo_2_2_2.Bar:5.5.5",       // backward compatible
	}
	fmt.Println("Server")
	fmt.Println(server)
	fmt.Println("Client | Compatible")
	for i := range clients {
		client, _ := ParseTypeCode(clients[i])
		fmt.Println(client, "|", server.IsCompatible(client))
	}
	// Output:
	// Server
	// foo.Abc_1_0_1.Foo_2_2_2.Bar:5.5.5
	// Client | Compatible
	// foo.Abc_1_0_1.Foo_2_2_2.Bar:5.5.5 | true
	// foo.Abc_1_0_1.Foo_2_2_2.Boo:5.5.5 | false
	// foo.Abc_1_0_1.Foo_2_2_2.Bar:6.5.5 | false
	// foo.Abc_1_0_1.Foo_2_2_2.Bar:4.5.5 | false
	// foo.Abc_1_0_1.Foo_2_2_2.Bar:5.6.5 | true
	// foo.Abc_1_0_1.Foo_2_2_2.Bar:5.4.5 | true
	// foo.Abc_1_0_1.Foo_2_2_2.Bar:5.5.6 | false
	// foo.Abc_1_0_1.Foo_2_2_2.Bar:5.5.4 | true
	// foo.Abc_1_0_1.Fuu_2_2_2.Bar:5.5.5 | false
	// foo.Abc_1_0_1.Foo_3_2_2.Bar:5.5.5 | false
	// foo.Abc_1_0_1.Foo_2_1_2.Bar:5.5.5 | true
	// foo.Abc_1_0_1.Foo_2_2_3.Bar:5.5.5 | false
	// foo.Abc_1_0_1.Foo_2_2_1.Bar:5.5.5 | true
	// foo.Abc.Foo_2_2_2.Bar:5.5.5 | true
}

func ExampleTypeCode_SegmentCodes() {
	code, _ := ParseTypeCode("foo.Foo_1_2_3.Bar:4.5.6")
	fmt.Println(code.SegmentCodes())
	// Output: [foo:1.0.0 Foo:1.2.3 Bar:4.5.6]
}
