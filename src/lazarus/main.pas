unit main;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, ActnList, StdCtrls,
  Menus;

type

  { TMainForm }

  TMainForm = class(TForm)
    MainMenu: TMainMenu;
    Memo1: TMemo;
    ConnectItem: TMenuItem;
  private

  public

  end;

var
  MainForm: TMainForm;

implementation

{$R *.lfm}

end.

