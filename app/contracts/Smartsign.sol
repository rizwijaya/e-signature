// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity ^0.8.0;
pragma abicoder v2;

contract Smartsign {
    uint256 documentCount;

    constructor() {
        documentCount = 1;
    }

    // ------Model data------ //
    //Model Document
    struct Document {
        bytes32 file; //Hash asli Bytes (Identifier)
        uint256 document_id; //ID document
        string creator_id;
        address creator; //Pembuat Documents
        string metadata;
        string hash_ori; //Hash Original
        string hash; //Hash akhir
        string ipfs; //data document di ipfs
        uint256 state; //1 Process Signed, 2 Signed
        uint256 mode; //mode ttd: 1 Permintaan ttd, 2 minta ttd (invite orang lain), 3 Single ttd
        uint256 createdtime; //tgl request ttd
        uint256 completedtime;
        bool exist; //check apakah dokument
        uint256 sign_reminder;
        mapping(address => Signers) signers; //Data alamat penandatangan
    }

    struct Signers {
        address sign_addr;
        uint256 sign_id;
        string signers_id; //id signatures
        string signers_hash;
        bool signers_state; //status ttd 
        uint sign_time; //tgl ttd
    }
    mapping(bytes32 => Document) documents;
    mapping(bytes32 => bytes32) signedDocs;
    //mapping(address => mapping(uint256))
    //------ End Model data ------//
    //------ Convertion Data ------//
    function stringToBytes32(string memory source) private pure returns (bytes32 result) {
        bytes memory tempEmptyStringTest = bytes(source);
        if (tempEmptyStringTest.length == 0) {
            return 0x0;
        }
        assembly {
                result := mload(add(source, 32))
        }
    }
    function bytes32ToString(bytes32 _bytes32) private pure returns (string memory) {
        bytes memory bytesArray = new bytes(32);
        for (uint256 i; i < 32; i++) {
            bytesArray[i] = _bytes32[i];
            }
        return string(bytesArray);
    }
    //------ End Convertion Data ------//
    //------ Signing Process ------//
    //Created Document
    function create(
        string memory _file, address creator, string memory _creator_id, 
        string memory _metadata, string memory _hash, 
        string memory _ipfs, uint256 _state, uint256 _mode, 
        uint256 _time, address[] memory _signers, string[] memory _signers_id
    ) public {
        bytes32 byte_id = stringToBytes32(_file);
        Document storage newDocument = documents[byte_id];
        newDocument.document_id = documentCount;
        newDocument.file = byte_id;
        newDocument.creator = creator;
        newDocument.creator_id = _creator_id;
        newDocument.metadata = _metadata;
        newDocument.hash_ori = _file;
        newDocument.hash = _hash;
        newDocument.ipfs = _ipfs;
        newDocument.state = _state;
        newDocument.mode = _mode;
        newDocument.createdtime = _time;
        newDocument.completedtime = _time;
        newDocument.exist = true;
        newDocument.sign_reminder = _signers.length;
        for (uint256 i=0; i<_signers.length; i++) {
            newDocument.signers[_signers[i]].sign_addr = _signers[i];
            newDocument.signers[_signers[i]].sign_id = i;
            newDocument.signers[_signers[i]].signers_id = _signers_id[i];
            newDocument.signers[_signers[i]].signers_hash = _hash;
            newDocument.signers[_signers[i]].signers_state = false;
            newDocument.signers[_signers[i]].sign_time = _time;
        }
        documentCount++;
    }
    //Get Document Data with Hash Original Files
    //Original hash document disimpan didalam file local
    function getDoc(string memory _file) public view returns(
        uint256, address, string memory, string memory, string memory, 
        string memory, string memory, uint256, uint256, uint256, uint256, bool
    ) {
        bytes32 byte_id = stringToBytes32(_file);
        Document storage temp = documents[byte_id];
        require(temp.exist == true, "Document not exist");
        return(temp.document_id, temp.creator, temp.creator_id, 
        temp.metadata, temp.hash_ori, temp.hash, temp.ipfs, temp.state,
        temp.mode, temp.createdtime, temp.completedtime, 
        temp.exist);
    }
    //Get Signatures Data in Documents
    function getSign(string memory _file, address _signers_id) public view returns(
        uint256, string memory, string memory, bool, uint
    ) {
        bytes32 byteFile = stringToBytes32(_file);
        Document storage temp = documents[byteFile];
        require(temp.exist == true, "Document not exist");
        require(temp.signers[_signers_id].sign_addr == _signers_id, "Signers not exist");
        return(temp.signers[_signers_id].sign_id, temp.signers[_signers_id].signers_id,
        temp.signers[_signers_id].signers_hash, temp.signers[_signers_id].signers_state, 
        temp.signers[_signers_id].sign_time);
    }
    //Signing Document dengan hash asli
    function signDoc(string memory _file, address _signers_id, string memory _signers_hash, string memory _ipfs, uint256 _time) public {
        bytes32 byteFile = stringToBytes32(_file); //Hash Original
        Document storage signDocument = documents[byteFile];
        require(signDocument.exist == true, "Document not exist");
        require(signDocument.signers[_signers_id].sign_time > 1, "Document not exist");
        require(signDocument.signers[_signers_id].signers_state == false, "You are signed this document");
        signDocument.sign_reminder--;
        signDocument.ipfs = _ipfs;
        signDocument.completedtime = _time;
        signDocument.hash = _signers_hash; //Update last hash
        signDocument.signers[_signers_id].signers_hash = _signers_hash; //users hash
        signDocument.signers[_signers_id].signers_state = true;
        signDocument.signers[_signers_id].sign_time = _time;
        bytes32 signedfile = stringToBytes32(_signers_hash);
        signedDocs[signedfile] = signDocument.file;
        if (signDocument.sign_reminder < 1) {
            signDocument.state = 2;
        }
    }
    //Get Documents Signed By Hash signed
    function getDocSigned(string memory _hash) public view returns(string memory) {
        bytes32 byteFile = stringToBytes32(_hash);
        bytes32 sign = signedDocs[byteFile];
        string memory temp = bytes32ToString(sign);
        return(temp);
    }
    //verification documents
    function verifyDoc(string memory _hash) public view returns(bool) {
        bytes32 signed = stringToBytes32(_hash);
        if (signedDocs[signed] == bytes32(0)) {
            return false;
        } else {
            return true;
        }
    }
     // ------ End Signing Process ------ //
}