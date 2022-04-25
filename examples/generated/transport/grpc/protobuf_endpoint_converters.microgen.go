// Code generated by microgen 0.9.0. DO NOT EDIT.

// Please, do not change functions names!
package transportgrpc

import (
	"context"
	"errors"
	transport "github.com/devimteam/microgen/examples/generated/transport"
	pb "github.com/devimteam/microgen/examples/protobuf"
	empty "github.com/golang/protobuf/ptypes/empty"
)

func _Encode_Uppercase_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil UppercaseRequest")
	}
	req := request.(*transport.UppercaseRequest)
	reqStringsMap, err := MapStringStringToProto(req.StringsMap)
	if err != nil {
		return nil, err
	}
	return &pb.UppercaseRequest{StringsMap: reqStringsMap}, nil
}

func _Encode_Count_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil CountRequest")
	}
	req := request.(*transport.CountRequest)
	return &pb.CountRequest{
		Symbol: req.Symbol,
		Text:   req.Text,
	}, nil
}

func _Encode_TestCase_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil TestCaseRequest")
	}
	req := request.(*transport.TestCaseRequest)
	reqComments, err := ListPtrCommentToProto(req.Comments)
	if err != nil {
		return nil, err
	}
	return &pb.TestCaseRequest{Comments: reqComments}, nil
}

func _Encode_DummyMethod_Request(ctx context.Context, request interface{}) (interface{}, error) {
	return &empty.Empty{}, nil
}

func _Encode_Uppercase_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil UppercaseResponse")
	}
	resp := response.(*transport.UppercaseResponse)
	return &pb.UppercaseResponse{Ans: resp.Ans}, nil
}

func _Encode_Count_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil CountResponse")
	}
	resp := response.(*transport.CountResponse)
	respPositions, err := ListIntToProto(resp.Positions)
	if err != nil {
		return nil, err
	}
	return &pb.CountResponse{
		Count:     int64(resp.Count),
		Positions: respPositions,
	}, nil
}

func _Encode_TestCase_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil TestCaseResponse")
	}
	resp := response.(*transport.TestCaseResponse)
	respTree, err := MapStringIntToProto(resp.Tree)
	if err != nil {
		return nil, err
	}
	return &pb.TestCaseResponse{Tree: respTree}, nil
}

func _Encode_DummyMethod_Response(ctx context.Context, response interface{}) (interface{}, error) {
	return &empty.Empty{}, nil
}

func _Decode_Uppercase_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil UppercaseRequest")
	}
	req := request.(*pb.UppercaseRequest)
	reqStringsMap, err := ProtoToMapStringString(req.StringsMap)
	if err != nil {
		return nil, err
	}
	return &transport.UppercaseRequest{StringsMap: reqStringsMap}, nil
}

func _Decode_Count_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil CountRequest")
	}
	req := request.(*pb.CountRequest)
	return &transport.CountRequest{
		Symbol: string(req.Symbol),
		Text:   string(req.Text),
	}, nil
}

func _Decode_TestCase_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil TestCaseRequest")
	}
	req := request.(*pb.TestCaseRequest)
	reqComments, err := ProtoToListPtrComment(req.Comments)
	if err != nil {
		return nil, err
	}
	return &transport.TestCaseRequest{Comments: reqComments}, nil
}

func _Decode_DummyMethod_Request(ctx context.Context, request interface{}) (interface{}, error) {
	return &empty.Empty{}, nil
}

func _Decode_Uppercase_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil UppercaseResponse")
	}
	resp := response.(*pb.UppercaseResponse)
	return &transport.UppercaseResponse{Ans: string(resp.Ans)}, nil
}

func _Decode_Count_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil CountResponse")
	}
	resp := response.(*pb.CountResponse)
	respPositions, err := ProtoToListInt(resp.Positions)
	if err != nil {
		return nil, err
	}
	return &transport.CountResponse{
		Count:     int(resp.Count),
		Positions: respPositions,
	}, nil
}

func _Decode_TestCase_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil TestCaseResponse")
	}
	resp := response.(*pb.TestCaseResponse)
	respTree, err := ProtoToMapStringInt(resp.Tree)
	if err != nil {
		return nil, err
	}
	return &transport.TestCaseResponse{Tree: respTree}, nil
}

func _Decode_DummyMethod_Response(ctx context.Context, response interface{}) (interface{}, error) {
	return &empty.Empty{}, nil
}
