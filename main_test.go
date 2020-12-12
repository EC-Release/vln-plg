/*
 * Copyright (c) 2016 General Electric Company. All rights reserved.
 *
 * The copyright to the computer software herein is the property of
 * General Electric Company. The software may be used and/or copied only
 * with the written permission of General Electric Company or in accordance
 * with the terms and conditions stipulated in the agreement/contract
 * under which the software has been supplied.
 *
 * author: apolo.yasuda@ge.com
 */

package main

import (
	"reflect"
	"testing"
)

func TestIPRoute_RegisterCidrList(t *testing.T) {
	type args struct {
		ips []string
	}
	tests := []struct {
		name    string
		i       *IPRoute
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &IPRoute{}
			if err := i.RegisterCidrList(tt.args.ips); (err != nil) != tt.wantErr {
				t.Errorf("IPRoute.RegisterCidrList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetVLANSetting(t *testing.T) {
	tests := []struct {
		name    string
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetVLANSetting()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVLANSetting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVLANSetting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
