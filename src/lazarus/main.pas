unit main;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, ActnList, StdCtrls,
  Menus, ComCtrls, Buttons, ExtCtrls, PairSplitter, Grids, ValEdit;

type

  { TMainForm }

  TMainForm = class(TForm)
    BtnFormat: TButton;
    Btn2: TButton;
    CBProtocols: TComboBox;
    CB2: TComboBox;
    MainMenu: TMainMenu;
    ConnectItem: TMenuItem;
    Memo1: TMemo;
    Memo2: TMemo;
    Panel1: TPanel;
    Panel2: TPanel;
    Panel3: TPanel;
    Splitter1: TSplitter;
    StringGrid1: TStringGrid;
    TestItem: TMenuItem;
  private

  public

  end;

var
  MainForm: TMainForm;

implementation

{$R *.lfm}

end.

