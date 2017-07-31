; Build the game with build.bat.
; Then compile this script with Inno Setup
; http://www.jrsoftware.org/isinfo.php
; to create a Windows installer for the game.

#define MyAppName "Breathless Parks"
#define MyAppPublisher "gonutz"
#define MyAppExeName "breathless_parks.exe"

[Setup]
; NOTE: The value of AppId uniquely identifies this application.
; Do not use the same AppId value in installers for other applications.
; (To generate a new GUID, click Tools | Generate GUID inside the IDE.)
AppId={{D58F25D6-BB05-4749-8C4E-8EBCF5577D53}
AppName={#MyAppName}
AppVerName=Breathless Parks
AppPublisher={#MyAppPublisher}
DefaultDirName={pf}\{#MyAppName}
DefaultGroupName={#MyAppName}
AllowNoIcons=yes
OutputDir=.\
OutputBaseFilename=Breathless_Parks_Setup
SetupIconFile=icon.ico
Compression=lzma
SolidCompression=yes

[Languages]
Name: "english"; MessagesFile: "compiler:Default.isl"

[Tasks]
Name: "desktopicon"; Description: "{cm:CreateDesktopIcon}"; GroupDescription: "{cm:AdditionalIcons}";
Name: "quicklaunchicon"; Description: "{cm:CreateQuickLaunchIcon}"; GroupDescription: "{cm:AdditionalIcons}"; Flags: unchecked; OnlyBelowVersion: 0,6.1

[Files]
Source: "breathless_parks.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "armchair.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "back_music.wav"; DestDir: "{app}"; Flags: ignoreversion
Source: "back_tiles.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "bird_left_down.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "bird_left_up.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "bird_right_down.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "bird_right_up.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "bird1.wav"; DestDir: "{app}"; Flags: ignoreversion
Source: "bird2.wav"; DestDir: "{app}"; Flags: ignoreversion
Source: "bird3.wav"; DestDir: "{app}"; Flags: ignoreversion
Source: "bird4.wav"; DestDir: "{app}"; Flags: ignoreversion
Source: "blink.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "closed_door.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "couch.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "couch_back.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "grays_anatomy.wav"; DestDir: "{app}"; Flags: ignoreversion
Source: "hit_table.wav"; DestDir: "{app}"; Flags: ignoreversion
Source: "instruction_left.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "instruction_right.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "matlock.wav"; DestDir: "{app}"; Flags: ignoreversion
Source: "nurse.wav"; DestDir: "{app}"; Flags: ignoreversion
Source: "nurse1.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "nurse2.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "nurse3.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "old_broad_talk_mouth.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "old_guy.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "old_guy_sitting.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "old_guy_sitting_blink.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "old_guy_win.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "old_guy1.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "old_guy2.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "old_guy3.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "other_dude_mouth_talks.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "other_dude_sitting.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "other_dude_winning.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "other_dude1.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "other_dude2.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "other_dude3.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "other_guy_wins.wav"; DestDir: "{app}"; Flags: ignoreversion
Source: "outside.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "painting1.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "painting2.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "shut_mouth.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "sleeping_woman.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "space_big.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "space_small.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "squeak.wav"; DestDir: "{app}"; Flags: ignoreversion
Source: "table.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "table_empty.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "tv.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "win.wav"; DestDir: "{app}"; Flags: ignoreversion
Source: "woman1.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "woman2.png"; DestDir: "{app}"; Flags: ignoreversion
Source: "woman3.png"; DestDir: "{app}"; Flags: ignoreversion
; NOTE: Don't use "Flags: ignoreversion" on any shared system files

[Icons]
Name: "{group}\{#MyAppName}"; Filename: "{app}\{#MyAppExeName}"
Name: "{group}\{cm:UninstallProgram,{#MyAppName}}"; Filename: "{uninstallexe}"
Name: "{commondesktop}\{#MyAppName}"; Filename: "{app}\{#MyAppExeName}"; Tasks: desktopicon
Name: "{userappdata}\Microsoft\Internet Explorer\Quick Launch\{#MyAppName}"; Filename: "{app}\{#MyAppExeName}"; Tasks: quicklaunchicon

[Run]
Filename: "{app}\{#MyAppExeName}"; Description: "{cm:LaunchProgram,{#StringChange(MyAppName, '&', '&&')}}"; Flags: nowait postinstall skipifsilent

