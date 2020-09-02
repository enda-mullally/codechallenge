;
; CodeChallenge InnoSetup Installer Script
;

#ifndef InstallerVersion
  #define InstallerVersion "1.0.0"
#endif

[Setup]
UninstallDisplayIcon={app}\MainIcon.ico
WizardStyle=modern
AppName=CodeChallenge
AppPublisher=EndaMullally
DefaultDirName={commonpf}\EndaMullally\CodeChallenge
DefaultGroupName=CodeChallenge
PrivilegesRequired=admin
AlwaysRestart=no
LicenseFile=
AppVersion={#InstallerVersion}
AppCopyright=Copyright 2020 Enda Mullally
DisableProgramGroupPage=true
DisableDirPage=true
Compression=lzma2
SetupIconFile=MainIcon.ico
WizardSmallImageFile=SetupIcon.bmp

AppVerName=CodeChallenge - Version: {#InstallerVersion}
OutputBaseFilename=setup

[Icons]
Name: "{group}\CodeChallenge"; Filename: "{app}\starwars.exe"

[Registry]
Root: HKLM; SubKey: "Software\EndaMullally"; ValueType: string; ValueName: ""; ValueData: ""; Flags: uninsdeletekey
Root: HKLM; SubKey: "Software\EndaMullally\CodeChallenge"; ValueType: string; ValueName: "App"; ValueData: "{app}\starwars.exe"; Flags: uninsdeletekey
Root: HKLM; SubKey: "Software\EndaMullally\CodeChallenge"; ValueType: string; ValueName: "Version"; ValueData: "{#InstallerVersion}"; Flags: uninsdeletekey

[Files]
Source: "..\starwars.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "MainIcon.ico"; DestDir: "{app}"; Flags: ignoreversion

[Run]
Filename: "{app}\starwars.exe"; Parameters:""; Description: ""; Flags: nowait shellexec skipifsilent

[UninstallRun]
