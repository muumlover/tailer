unit connect;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, StdCtrls, ExtCtrls,
  ComboEx;

type

  { TConnectForm }

  TConnectForm = class(TForm)
    ComboBox1: TComboBox;
    ComboBoxEx1: TComboBoxEx;
    ListBox1: TListBox;
    Panel1: TPanel;
  private

  public

  end;

var
  ConnectForm: TConnectForm;

implementation

{$R *.lfm}

end.

