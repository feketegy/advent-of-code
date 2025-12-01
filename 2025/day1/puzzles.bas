Call puzzle1
Call puzzle2

Sub puzzle1 ()
    Dim fileLine As String

    Dim boundary As Integer
    boundary = 100

    Dim cursor As Integer
    cursor = 50

    Dim counter As Integer


    Open "input.txt" For Input As #1
    Do While Not EOF(1)
        Line Input #1, fileLine

        Dim dir As String
        Dim steps As Integer

        dir = Left$(fileLine, 1)
        steps = Val(Mid$(fileLine, 2))

        If dir = "L" Then
            steps = steps * -1
        End If

        cursor = (cursor + steps) Mod boundary
        If cursor < 0 Then
            cursor = cursor + boundary
        End If

        If cursor = 0 Then
            counter = counter + 1
        End If
    Loop

    Close #1

    Print "Puzzle 1 - password: "; counter
End Sub


Sub puzzle2 ()
    Dim fileLine As String

    Dim boundary As Integer
    boundary = 100

    Dim cursor As Integer
    cursor = 50

    Dim counter As Integer


    Open "input.txt" For Input As #1
    Do While Not EOF(1)
        Line Input #1, fileLine

        Dim dir As String
        Dim steps As Integer
        Dim stepper As Integer

        dir = Left$(fileLine, 1)
        steps = Val(Mid$(fileLine, 2))

        stepper = 1

        If dir = "L" Then
            stepper = -1
        End If

        For i = 1 To steps Step 1
            cursor = cursor + stepper
            If cursor = boundary Then
                cursor = 0
            Else
                If cursor < 0 Then
                    cursor = boundary - 1
                End If
            End If

            If cursor = 0 Then
                counter = counter + 1
            End If
        Next i
    Loop

    Close #1

    Print "Puzzle 2 - password: "; counter
End Sub
