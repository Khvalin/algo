package playlistlongestinterval

import "testing"

func Test_playlistLongestInterval(t *testing.T) {
	type args struct {
		songs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		/*	{
				name: "it runs",
				args: args{[]string{}},
				want: 0,
			},
			{
				name: "sample test 1",
				args: args{[]string{"A (1:30)",
					"B (2:00)",
					"A (1:30)",
					"C (1:00)"}},
				want: 270,
			},
			{
				name: "sample test 3",
				args: args{[]string{
					"A (1:30)",
					"C (1:00)",
					"C (1:00)",
					"B (2:00)",
					"B (2:00)",
					"A (1:30)",
				}},
				want: 210,
			},*/

		{
			name: "sample test 7",
			args: args{},
			want: 1916,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := playlistLongestInterval(tt.args.songs); got != tt.want {
				t.Errorf("playlistLongestInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSample7(t *testing.T) {
	expected := 1916
	input := []string{
		"HEOIG (1:52)",
		"F (9:24)",
		"IXDK (0:42)",
		"F (9:24)",
		"D (2:11)",
		"HEOIG (1:52)",
		"IXDK (0:42)",
		"GEAA (2:19)",
		"D (2:11)",
		"IDNQ (9:10)",
		"VNWBLVNUEZ (0:13)",
		"UHHZILNA (9:47)",
		"UZVZ (5:42)",
		"IXDK (0:42)",
		"VNWBLVNUEZ (0:13)",
		"LY (2:48)",
		"UZVZ (5:42)",
		"IDNQ (9:10)",
		"G (3:02)",
		"G (3:02)",
		"IYW (4:26)",
		"UHHZILNA (9:47)",
		"E (4:05)",
		"QNYZXPC (0:59)",
		"UZVZ (5:42)",
	}
	result := playlistLongestInterval(input)

	if result == expected {
		t.Log("passed", result)
	} else {
		t.Error("failed! got ", result, " instead of ", 1916)
	}
}
